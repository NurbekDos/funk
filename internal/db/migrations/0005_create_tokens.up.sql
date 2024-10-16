    CREATE TABLE tokens (
        id SERIAL PRIMARY KEY,
        case_id INT REFERENCES cases(id) ON DELETE CASCADE,
        type VARCHAR(50) NOT NULL,
        symbol VARCHAR(10) NOT NULL,
        name VARCHAR(64) NOT NULL,
        price NUMERIC(10, 2) NOT NULL,
        issuer_number INT NOT NULL,
        image BYTEA,
        -- company_name VARCHAR(255) NOT NULL,
        company_area VARCHAR(64) NOT NULL,
        company_capital NUMERIC(15, 2),
        description TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT NOW(),
        deleted_at TIMESTAMPTZ
    );
