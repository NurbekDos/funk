CREATE TABLE sto_config (
    id SERIAL PRIMARY KEY,
    start_datetime TIMESTAMP NOT NULL,
    end_datetime TIMESTAMP,
    total_aiming_price DECIMAL(15, 2) NOT NULL,
    max_sales_number INT,
    expected_annual_interest_rate DECIMAL(5, 2),
    token_receival_timing_id INT REFERENCES token_receival_timing(id) ON DELETE CASCADE,
    yield_description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
