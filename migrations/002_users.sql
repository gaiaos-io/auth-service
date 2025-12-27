CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

	contact_email CITEXT UNIQUE,
	status SMALLINT NOT NULL,
	
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

	CHECK (status in (1, 2, 3, 4))
);

-- #############
-- Update guard
-- #############

CREATE OR REPLACE FUNCTION users_update_guard()
RETURNS TRIGGER AS $$
BEGIN
	IF
		NEW.id			IS DISTINCT FROM OLD.id OR
        NEW.created_at  IS DISTINCT FROM OLD.created_at OR
		NEW.updated_at	IS DISTINCT FROM OLD.updated_at
    THEN
        RAISE EXCEPTION
            'Only contact_email and status may be updated on users';
    END IF;

	NEW.updated_at = now();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS users_update_guard ON users;

CREATE TRIGGER users_update_guard
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION users_update_guard();
