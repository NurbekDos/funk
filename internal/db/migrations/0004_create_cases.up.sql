CREATE TABLE cases (
    id SERIAL PRIMARY KEY,
    case_name VARCHAR(64) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);