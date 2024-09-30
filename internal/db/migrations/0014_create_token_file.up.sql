CREATE TABLE token_file (
    id SERIAL PRIMARY KEY,
    token_id INT REFERENCES token(id) ON DELETE CASCADE, 
    title VARCHAR(64),
    sort INT,
    file_path VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
