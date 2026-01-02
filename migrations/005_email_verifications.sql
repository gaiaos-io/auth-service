-- ============================================================================
-- Enums
-- ============================================================================

DO $$
BEGIN
    CREATE TYPE email_purpose AS ENUM (
		'auth',
		'contact'
	);
EXCEPTION
    WHEN duplicate_object THEN NULL;
END$$;

ALTER TYPE email_purpose ADD VALUE IF NOT EXISTS 'auth';
ALTER TYPE email_purpose ADD VALUE IF NOT EXISTS 'contact';

-- ============================================================================
-- Email verifications table
-- ============================================================================

CREATE TABLE IF NOT EXISTS email_verifications (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

	-- Ownership
	account_id UUID NULL REFERENCES accounts(id) ON DELETE CASCADE,
	auth_identity_id UUID NULL REFERENCES auth_identities(id) ON DELETE CASCADE,

	-- Email context
	email CITEXT NOT NULL,
	purpose email_purpose NOT NULL,

	-- Token lifecycle
	token_hash BYTEA NOT NULL UNIQUE,
	expires_at TIMESTAMPTZ NOT NULL,
	consumed_at TIMESTAMPTZ,

	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

	-- ------------------------------------------------------------------------
    -- Constraints
    -- ------------------------------------------------------------------------

    CONSTRAINT email_verifications_owner_matches_purpose
        CHECK (
            (purpose = 'auth'
                AND auth_identity_id IS NOT NULL
                AND account_id IS NULL)
            OR
            (purpose = 'contact'
                AND account_id IS NOT NULL
                AND auth_identity_id IS NULL)
        ),

    CONSTRAINT email_verifications_consumed_after_creation
        CHECK (
            consumed_at IS NULL
            OR consumed_at >= created_at
        ),
	
	CONSTRAINT email_verifications_consumed_before_expiry
        CHECK (
            consumed_at IS NULL
            OR consumed_at <= expires_at
        ),

    CONSTRAINT email_verifications_expires_after_creation
        CHECK (expires_at > created_at),

	CONSTRAINT email_verifications_token_hash_length
        CHECK (octet_length(token_hash) = 32)

);

-- ============================================================================
-- Update guard
-- ============================================================================

CREATE OR REPLACE FUNCTION email_verifications_update_guard_fn()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.consumed_at IS NOT NULL OR OLD.expires_at <= NOW() THEN
        RAISE EXCEPTION
            'email_verifications rows cannot be updated after consumption or expiry';
    END IF;	

	-- Immutable fields	
	IF
		NEW.id					IS DISTINCT FROM OLD.id OR
        NEW.account_id     		IS DISTINCT FROM OLD.account_id OR
		NEW.auth_identity_id	IS DISTINCT FROM OLD.auth_identity_id OR
        NEW.email       		IS DISTINCT FROM OLD.email OR
		NEW.email				IS DISTINCT FROM OLD.purpose OR
        NEW.token_hash  		IS DISTINCT FROM OLD.token_hash OR
        NEW.expires_at  		IS DISTINCT FROM OLD.expires_at OR
        NEW.created_at  		IS DISTINCT FROM OLD.created_at
    THEN
        RAISE EXCEPTION
            'Only consumed_at may be updated on email_verifications';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS email_verifications_update_guard_trg
ON email_verifications;

CREATE TRIGGER email_verifications_update_guard_trg
BEFORE UPDATE ON email_verifications
FOR EACH ROW
EXECUTE FUNCTION email_verifications_update_guard_fn();

-- ============================================================================
-- Indexes
-- ============================================================================

-- Expiration sweeps / cleanup jobs
CREATE INDEX IF NOT EXISTS email_verifications_expires_unconsumed_idx
ON email_verifications (expires_at)
WHERE consumed_at IS NULL;

-- Pending auth-identity verifications
CREATE INDEX IF NOT EXISTS email_verifications_auth_identity_unconsumed_idx
ON email_verifications (auth_identity_id)
WHERE purpose = 'auth'
  AND consumed_at IS NULL;

-- Pending account email verifications
CREATE INDEX IF NOT EXISTS email_verifications_account_unconsumed_idx
ON email_verifications (account_id)
WHERE purpose = 'contact'
  AND consumed_at IS NULL;
