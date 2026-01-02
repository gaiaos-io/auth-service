-- ============================================================================
-- Password credentials table
-- ============================================================================

CREATE TABLE IF NOT EXISTS password_credentials (
	auth_identity_id UUID PRIMARY KEY REFERENCES auth_identities(id) ON DELETE CASCADE,
	
	password_hash TEXT NOT NULL,	

	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	
	CONSTRAINT password_credentials_argon2id_hash
		CHECK (password_hash LIKE '$argon2id$%')
);

-- ============================================================================
-- Update guard
-- ============================================================================

CREATE OR REPLACE FUNCTION password_credentials_update_guard_fn()
RETURNS TRIGGER AS $$
BEGIN
	-- Immutable fields
	IF
		NEW.auth_identity_id	IS DISTINCT FROM OLD.auth_identity_id OR
        NEW.created_at  		IS DISTINCT FROM OLD.created_at OR
	THEN
        RAISE EXCEPTION
            'Only password_hash may be updated on password_credentials';
    END IF;

	NEW.updated_at = NOW();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS password_credentials_update_guard_trg
ON password_credentials;

CREATE TRIGGER password_credentials_update_guard_trg
BEFORE UPDATE ON password_credentials
FOR EACH ROW
EXECUTE FUNCTION password_credentials_update_guard_fn()
