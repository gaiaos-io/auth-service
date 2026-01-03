-- ============================================================================
-- Sessions table
-- ============================================================================

CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,

    -- Refresh token rotation
    refresh_token_hash BYTEA NOT NULL UNIQUE,
    previous_refresh_token_hash BYTEA,

    -- Session activity & lifecycle
    last_used_at TIMESTAMPTZ,
    rotated_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ,

    -- Device metadata
    account_agent TEXT,
    ip_address INET,
    device_label TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    -- ------------------------------------------------------------------------
    -- Constraints
    -- ------------------------------------------------------------------------

	CONSTRAINT sessions_created_before_expires
    	CHECK (created_at < expires_at),

	CONSTRAINT sessions_revoked_after_creation
    	CHECK (revoked_at IS NULL OR created_at <= revoked_at),

	CONSTRAINT sessions_refresh_token_length
    	CHECK (octet_length(refresh_token_hash) = 32),

	CONSTRAINT sessions_previous_refresh_token_length
    	CHECK (
        	previous_refresh_token_hash IS NULL
        	OR octet_length(previous_refresh_token_hash) = 32
    	),

	CONSTRAINT sessions_refresh_tokens_must_differ
    	CHECK (
        	previous_refresh_token_hash IS NULL
        	OR previous_refresh_token_hash <> refresh_token_hash
    	),
	
	CONSTRAINT sessions_max_lifetime_90d
    	CHECK (expires_at <= created_at + INTERVAL '90 days')
);

-- ============================================================================
-- Update guard
-- ============================================================================

CREATE OR REPLACE FUNCTION sessions_update_guard_fn()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.revoked_at IS NOT NULL OR OLD.expires_at <= NOW() THEN
        RAISE EXCEPTION
            'session rows cannot be updated after revocation or expiration';
    END IF;
	
    IF OLD.rotated_at IS NOT NULL
    AND NEW.rotated_at IS NOT NULL
    AND NEW.rotated_at < OLD.rotated_at THEN
        RAISE EXCEPTION
            'rotated_at must move forward';
    END IF;

    IF NEW.refresh_token_hash IS DISTINCT FROM OLD.refresh_token_hash
    AND NEW.rotated_at IS NULL THEN
        RAISE EXCEPTION
            'rotated_at must be set when refresh_token_hash changes';
    END IF;
	
    IF OLD.rotated_at IS NOT NULL
    AND NEW.rotated_at IS NULL THEN
        RAISE EXCEPTION
            'rotated_at cannot be set to NULL once set';
    END IF;
	
    IF NEW.expires_at IS DISTINCT FROM OLD.expires_at THEN
        IF NEW.refresh_token_hash IS NOT DISTINCT FROM OLD.refresh_token_hash THEN
            RAISE EXCEPTION
                'expires_at may only be updated during refresh token rotation';
        END IF;

        IF NEW.expires_at <= OLD.expires_at THEN
            RAISE EXCEPTION
                'expires_at must move forward';
        END IF;
    END IF;
	
	-- Immutable fields
    IF
        NEW.id                   IS DISTINCT FROM OLD.id OR
        NEW.account_id              IS DISTINCT FROM OLD.account_id OR
        NEW.created_at           IS DISTINCT FROM OLD.created_at
    THEN
        RAISE EXCEPTION
			'only refresh token rotation, revocation, usage timestamps, and device metadata may be updated on sessions';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS sessions_update_guard_trg
ON sessions;

CREATE TRIGGER sessions_update_guard_trg
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE FUNCTION sessions_update_guard_fn();

-- ============================================================================
-- Indexes
-- ============================================================================

-- Active sessions per account
CREATE INDEX IF NOT EXISTS sessions_account_active_idx
ON sessions(account_id)
WHERE revoked_at IS NULL;

-- Expiration sweeps / cleanup jobs
CREATE INDEX IF NOT EXISTS sessions_expires_active_idx
ON sessions(expires_at)
WHERE revoked_at IS NULL;

-- Refresh token lookup
CREATE UNIQUE INDEX IF NOT EXISTS sessions_refresh_token_active_uidx
ON sessions(refresh_token_hash)
WHERE revoked_at IS NULL;

-- Replay detection lookup
CREATE INDEX IF NOT EXISTS sessions_previous_refresh_token_active_idx
ON sessions(previous_refresh_token_hash)
WHERE revoked_at IS NULL;
