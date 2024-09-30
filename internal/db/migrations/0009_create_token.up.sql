CREATE TABLE token (
    id SERIAL PRIMARY KEY,
    case_id INT REFERENCES case(id) ON DELETE CASCADE,
    token_type_id INT REFERENCES token_type(id) ON DELETE CASCADE,
    token_symbol VARCHAR(20) NOT NULL,
    token_name VARCHAR(64) NOT NULL,
    token_price DECIMAL(10, 2) NOT NULL,
    issue_number VARCHAR(50) NOT NULL,
    token_image_file VARCHAR(255),--TODO image file
    company_name VARCHAR(64) NOT NULL,
    company_area VARCHAR(64) NOT NULL,
    company_capital DECIMAL(15, 2),
    description TEXT
);
