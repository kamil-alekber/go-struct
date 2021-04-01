CREATE TABLE IF NOT EXISTS migrations (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR (50) UNIQUE NOT NULL,
    tag VARCHAR (50),
    body VARCHAR NOT NULL,
    migration_type VARCHAR (50) NOT NULL,
    migrated_at TIMESTAMP NOT NULL
)