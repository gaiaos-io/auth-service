CREATE TABLE IF NOT EXISTS password_credentials (
	auth_identity_id UUID PRIMARY KEY REFERENCES auth_identities(id) ON DELETE CASCADE,
	
	password_hash VARCHAR(255) NOT NULL,	

	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

	CHECK (
    	password_hash LIKE '$argon2id$%' AND
    	CHAR_LENGTH(password_hash) >= 90
	)
);

-- #############
-- Update guard
-- #############

CREATE OR REPLACE FUNCTION password_credentials_update_guard()
RETURNS TRIGGER AS $$
BEGIN
	IF
		NEW.id			IS DISTINCT FROM OLD.id OR
        NEW.created_at  IS DISTINCT FROM OLD.created_at OR
		NEW.updated_at	IS DISTINCT FROM OLD.updated_at
	THEN
        RAISE EXCEPTION
            'Only password_hash may be updated on password_credentials';
    END IF;

	NEW.updated_at = now();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS password_credentials_update_guard ON password_credentials;

CREATE TRIGGER password_credentials_update_guard
BEFORE UPDATE ON password_credentials
FOR EACH ROW
EXECUTE FUNCTION password_credentials_update_guard()
