CREATE TABLE IF NOT EXISTS refresh_tokens (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

	token_hash TEXT NOT NULL UNIQUE,
	revoked_at TIMESTAMPTZ,
	expires_at TIMESTAMPTZ NOT NULL,

	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

	CHECK (created_at < expires_at),
	CHECK (revoked_at IS NULL OR created_at <= revoked_at),

	CHECK (length(token_hash) >= 32)
);

-- #############
-- Update guard
-- #############

CREATE OR REPLACE FUNCTION email_verifications_update_guard()
RETURNS TRIGGER AS $$
BEGIN
	IF OLD.revoked_at IS NOT NULL OR OLD.expires_at <= now() THEN
        RAISE EXCEPTION
            'refresh_tokens cannot be updated after revocation or expiration';
    END IF;
	
	IF
		NEW.id			IS DISTINCT FROM OLD.id OR
        NEW.user_id     IS DISTINCT FROM OLD.user_id OR
        NEW.token_hash  IS DISTINCT FROM OLD.token_hash OR
        NEW.expires_at  IS DISTINCT FROM OLD.expires_at OR
        NEW.created_at  IS DISTINCT FROM OLD.created_at
    THEN
        RAISE EXCEPTION
            'Only revoked_at may be updated on refresh_tokens';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS refresh_tokens_update_guard ON refresh_tokens;

CREATE TRIGGER refresh_tokens_update_guard
BEFORE UPDATE ON refresh_tokens
FOR EACH ROW
EXECUTE FUNCTION refresh_tokens_update_guard();

-- ########
-- Indexes
-- ########

CREATE INDEX IF NOT EXISTS refresh_tokens_user_idx
ON refresh_tokens(user_id)
WHERE revoked_at IS NULL;

CREATE INDEX IF NOT EXISTS refresh_tokens_expires_at_idx
ON refresh_tokens(expires_at)
WHERE revoked_at IS NULL;
