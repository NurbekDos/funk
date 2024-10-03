CREATE TABLE issuer (
    id SERIAL PRIMARY KEY,
    email VARCHAR(64) NOT NULL,
    phone_number VARCHAR(20),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    password VARCHAR(128),
    created_by INTEGER REFERENCES admin(id)
);
