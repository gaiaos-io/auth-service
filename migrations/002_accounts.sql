-- ============================================================================
-- Enums
-- ============================================================================

DO $$
BEGIN
    CREATE TYPE account_status AS ENUM (
		'unverified',	
		'active',
		'disabled'
	);
EXCEPTION
    WHEN duplicate_object THEN NULL;
END$$;

ALTER TYPE account_status ADD VALUE IF NOT EXISTS 'unverified';
ALTER TYPE account_status ADD VALUE IF NOT EXISTS 'active';
ALTER TYPE account_status ADD VALUE IF NOT EXISTS 'disabled';

-- ============================================================================
-- Accounts table
-- ============================================================================

CREATE TABLE IF NOT EXISTS accounts (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

	contact_email CITEXT UNIQUE,
	status account_status NOT NULL,
	
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);

-- ============================================================================
-- Update guard
-- ============================================================================

CREATE OR REPLACE FUNCTION accounts_update_guard_fn()
RETURNS TRIGGER AS $$
BEGIN
	-- Immutable fields
	IF
		NEW.id			IS DISTINCT FROM OLD.id OR
        NEW.created_at  IS DISTINCT FROM OLD.created_at OR
    THEN
        RAISE EXCEPTION
            'Only contact_email and status may be updated on accounts';
    END IF;

	NEW.updated_at = NOW();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS accounts_update_guard_trg
ON accounts;

CREATE TRIGGER accounts_update_guard_trg
BEFORE UPDATE ON accounts
FOR EACH ROW
EXECUTE FUNCTION accounts_update_guard_fn();
