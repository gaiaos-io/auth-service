-- ============================================================================
-- Enums
-- ============================================================================

DO $$
BEGIN
    CREATE TYPE auth_provider AS ENUM (
		'email_password',
		'google',
		'github'
	);
EXCEPTION
    WHEN duplicate_object THEN NULL;
END$$;

ALTER TYPE auth_provider ADD VALUE IF NOT EXISTS 'email_password';
ALTER TYPE auth_provider ADD VALUE IF NOT EXISTS 'google';
ALTER TYPE auth_provider ADD VALUE IF NOT EXISTS 'github';

-- ============================================================================
-- Authentication identities table
-- ============================================================================

CREATE TABLE IF NOT EXISTS auth_identities (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	
	-- Ownership
	account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,

	-- Provider idenitity
	provider SMALLINT NOT NULL,
	provider_subject TEXT NOT NULL,
	provider_email CITEXT,

	-- Verification
	verified_at TIMESTAMPTZ,
	
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	
	-- ------------------------------------------------------------------------
    -- Constraints
    -- ------------------------------------------------------------------------

    CONSTRAINT auth_identities_unique_account_provider
        UNIQUE (account_id, provider),

    CONSTRAINT auth_identities_unique_provider_subject
        UNIQUE (provider, provider_subject),

    CONSTRAINT auth_identities_email_password_subject
        CHECK (
            provider <> 'email_password'
            OR (
                provider_email IS NOT NULL
                AND provider_subject = provider_email
            )
        ),

    CONSTRAINT auth_identities_oauth_always_verified
        CHECK (
            provider = 'email_password'
            OR verified_at IS NOT NULL
        )
);

-- ============================================================================
-- Update guard
-- ============================================================================

CREATE OR REPLACE FUNCTION auth_identities_update_guard_fn()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.verified_at IS NOT NULL
       AND NEW.verified_at IS DISTINCT FROM OLD.verified_at THEN
        RAISE EXCEPTION
            'verified_at cannot be modified once set';
    END IF;
	
	-- Immutable fields
	IF
		NEW.id					IS DISTINCT FROM OLD.id OR
		NEW.account_id				IS DISTINCT FROM OLD.account_id OR
		NEW.provider        	IS DISTINCT FROM OLD.provider OR
        NEW.provider_subject	IS DISTINCT FROM OLD.provider_subject OR
        NEW.created_at  		IS DISTINCT FROM OLD.created_at OR
    THEN
        RAISE EXCEPTION
            'Only provider_email and verified_at may be updated on auth_identities';
    END IF;	

	NEW.updated_at = NOW();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS auth_identities_update_guard_trg
ON auth_identities;

CREATE TRIGGER auth_identities_update_guard_trg
BEFORE UPDATE ON auth_identities
FOR EACH ROW
EXECUTE FUNCTION auth_identities_update_guard_fn()

-- ============================================================================
-- Indexes
-- ============================================================================

-- Verified identities per account
CREATE INDEX IF NOT EXISTS auth_identities_account_verified_idx
ON auth_identities (account_id)
WHERE verified_at IS NOT NULL;
