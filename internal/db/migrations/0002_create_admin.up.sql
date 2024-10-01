CREATE TABLE admin (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    role VARCHAR(32) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    password VARCHAR(128)
);