CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    email VARCHAR(64) UNIQUE NOT NULL,
    phone_number VARCHAR(20),
    first_name VARCHAR(32),--personal_name
    last_name VARCHAR(32),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    password VARCHAR(128),
    email_verified_at TIMESTAMPTZ
);