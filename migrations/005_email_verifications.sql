CREATE TABLE IF NOT EXISTS email_verifications (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	email CITEXT NOT NULL,

	token_hash TEXT NOT NULL UNIQUE,
	expires_at TIMESTAMPTZ NOT NULL,
	used_at TIMESTAMPTZ,

	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

	CHECK (expires_at > used_at),
	CHECK (used_at IS NULL OR used_at >= created_at)
);

-- #############
-- Update guard
-- #############

CREATE OR REPLACE FUNCTION email_verifications_update_guard()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.used_at IS NOT NULL THEN
        RAISE EXCEPTION
            'email_verifications rows cannot be updated after being used';
    END IF;
	
	IF
		NEW.id			IS DISTINCT FROM OLD.id OR
        NEW.user_id     IS DISTINCT FROM OLD.user_id OR
        NEW.email       IS DISTINCT FROM OLD.email OR
        NEW.token_hash  IS DISTINCT FROM OLD.token_hash OR
        NEW.expires_at  IS DISTINCT FROM OLD.expires_at OR
        NEW.created_at  IS DISTINCT FROM OLD.created_at
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

CREATE INDEX IF NOT EXISTS email_verifications_user_id_idx
ON email_verifications (user_id);

CREATE INDEX IF NOT EXISTS email_verifications_expires_at_idx
ON email_verifications (expires_at)
WHERE used_at IS NULL;
