### Run command for docker

-> docker exec -it 3b1b398272ea psql -U user -d student

CREATE FUNCTION moveJob() RETURNS void AS $BODY$ BEGIN DO $$ BEGIN IF NOT EXISTS (
SELECT 1
FROM pg_type
WHERE typname = 'migration_type'
) THEN CREATE TYPE migration_type AS ENUM ('UP', 'DOWN');
END IF;
END $$;
CREATE TABLE IF NOT EXISTS migrations (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR (50) UNIQUE NOT NULL,
    tag VARCHAR (50),
    body VARCHAR NOT NULL,
    migration_type migration_type,
    migrated_at TIMESTAMP NOT NULL
)
END;
$BODY$ LANGUAGE plpgsql VOLATILE SECURITY DEFINER COST 100;
select moveJob($1)
