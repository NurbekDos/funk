CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(64) NOT NULL,
    first_name VARCHAR(32),
    last_name VARCHAR(32),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    password VARCHAR(128),
    email_verified_at TIMESTAMPTZ
);