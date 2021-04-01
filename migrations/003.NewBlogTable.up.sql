CREATE TABLE IF NOT EXISTS blogs (
    id UUID PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL,
    title VARCHAR (255) NOT NULL,
    slug VARCHAR (255) NOT NULL,
    body VARCHAR,
    created TIMESTAMP NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
)