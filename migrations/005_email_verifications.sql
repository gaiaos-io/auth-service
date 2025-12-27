CREATE TABLE IF NOT EXISTS email_verifications (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

	user_id UUID NULL REFERENCES users(id) ON DELETE CASCADE,
	auth_identity_id UUID NULL REFERENCES auth_identities(id) ON DELETE CASCADE,

	email CITEXT NOT NULL,
	email_purpose SMALLINT NOT NULL,

	token_hash TEXT NOT NULL UNIQUE,
	expires_at TIMESTAMPTZ NOT NULL,
	used_at TIMESTAMPTZ,

	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

	CHECK (email_purpose in (1, 2))
	CHECK (
		(email_purpose = 1
			AND (auth_identity_id IS NOT NULL and user_id is NULL))
		OR
		(email_purpose = 2
			AND (user_id IS NOT NULL AND auth_identity_id IS NULL))
	)
	CHECK (expires_at > used_at),
	CHECK (used_at IS NULL OR used_at >= created_at)
);

-- #############
-- Update guard
-- #############

CREATE OR REPLACE FUNCTION email_verifications_update_guard()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.used_at IS NOT NULL OR OLD.expires_at <= now() THEN
        RAISE EXCEPTION
            'email_verifications rows cannot be updated after being used or having expired';
    END IF;
	
	IF
		NEW.id					IS DISTINCT FROM OLD.id OR
        NEW.user_id     		IS DISTINCT FROM OLD.user_id OR
		NEW.auth_identity_id	IS DISTINCT FROM OLD.auth_identity_id OR
        NEW.email       		IS DISTINCT FROM OLD.email OR
		NEW.email_purpose		IS DISTINCT FROM OLD.email_purpose OR
        NEW.token_hash  		IS DISTINCT FROM OLD.token_hash OR
        NEW.expires_at  		IS DISTINCT FROM OLD.expires_at OR
        NEW.created_at  		IS DISTINCT FROM OLD.created_at
    THEN
        RAISE EXCEPTION
            'Only used_at may be updated on email_verifications';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS email_verifications_update_guard ON email_verifications;

CREATE TRIGGER email_verifications_update_guard
BEFORE UPDATE ON email_verifications
FOR EACH ROW
EXECUTE FUNCTION email_verifications_update_guard();

-- ########
-- Indexes
-- ########

CREATE INDEX IF NOT EXISTS email_verifications_expires_at_idx
ON email_verifications (expires_at)
WHERE consumed_at IS NULL;

CREATE INDEX IF NOT EXISTS email_verifications_auth_identity_idx
ON email_verifications (auth_identity_id)
WHERE email_purpose = 1
	AND consumed_at IS NULL

CREATE INDEX IF NOT EXISTS email_verifications_user_idx
ON email_verifications (user_id)
WHERE email_purpose = 2
	AND consumed_at IS NULL
