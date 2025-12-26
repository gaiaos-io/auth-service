CREATE TABLE IF NOT EXISTS auth_identities (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

	provider CITEXT NOT NULL,
	provider_user_id TEXT NOT NULL,
	provider_email CITEXT,
	
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

	UNIQUE (provider, provider_user_id),
	UNIQUE (user_id, provider)
);

-- #############
-- Update guard
-- #############

CREATE OR REPLACE FUNCTION auth_identities_update_guard()
RETURNS TRIGGER AS $$
BEGIN
	IF
		NEW.id					IS DISTINCT FROM OLD.id OR
		NEW.user_id				IS DISTINCT FROM OLD.user_id OR
		NEW.provider        	IS DISTINCT FROM OLD.provider OR
        NEW.provider_user_id	IS DISTINCT FROM OLD.provider_user_id OR
        NEW.created_at  		IS DISTINCT FROM OLD.created_at OR
		NEW.updated_at			IS DISTINCT FROM OLD.updated_at
    THEN
        RAISE EXCEPTION
            'Only provider_email may be updated on auth_identities';
    END IF;

	NEW.updated_at = now();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS auth_identities_update_guard ON auth_identities;

CREATE TRIGGER auth_identities_update_guard
BEFORE UPDATE ON auth_identities
FOR EACH ROW
EXECUTE FUNCTION auth_identities_update_guard()

-- ########
-- Indexes
-- ########

CREATE INDEX IF NOT EXISTS auth_identities_user_id_idx
ON auth_identities (user_id);
