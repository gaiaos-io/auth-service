CREATE TABLE IF NOT EXISTS auth_identities (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

	provider SMALLINT NOT NULL,
	provider_subject TEXT NOT NULL,
	provider_email CITEXT,

	verified_at TIMESTAMPTZ,
	
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

	UNIQUE (provider, provider_subject),
	UNIQUE (user_id, provider),
	CHECK (provider in (0, 1, 2))
	CHECK (
    	(provider != 0 AND
			(verified_at IS NOT NULL AND verified_at = created_at)
		) OR
    	(provider_email IS NOT NULL AND provider_subject = provider_email)
	)
);

-- #############
-- Update guard
-- #############

CREATE OR REPLACE FUNCTION auth_identities_update_guard()
RETURNS TRIGGER AS $$
BEGIN
	IF OLD.verified_at IS NOT NULL AND NEW.verified_at IS DISTINCT FROM OLD.verified_at THEN
		RAISE EXCEPTION
			'Once set, verified_at cannot be modified';
	END IF;
	
	IF
		NEW.id					IS DISTINCT FROM OLD.id OR
		NEW.user_id				IS DISTINCT FROM OLD.user_id OR
		NEW.provider        	IS DISTINCT FROM OLD.provider OR
        NEW.provider_subject	IS DISTINCT FROM OLD.provider_subject OR
        NEW.created_at  		IS DISTINCT FROM OLD.created_at OR
		NEW.updated_at			IS DISTINCT FROM OLD.updated_at
    THEN
        RAISE EXCEPTION
            'Only provider_email and verified_at may be updated on auth_identities';
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

CREATE INDEX IF NOT EXISTS auth_identities_verified_user_id_idx
ON auth_identities (user_id)
WHERE verified_at IS NOT NULL;

CREATE INDEX IF NOT EXISTS auth_identities_unverified_created_at_idx
ON auth_identities (created_at)
WHERE verified_at IS NULL;
