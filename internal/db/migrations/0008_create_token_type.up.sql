CREATE TABLE token_type (
    id SERIAL PRIMARY KEY,
    token_type_name VARCHAR(32) UNIQUE NOT NULL
);