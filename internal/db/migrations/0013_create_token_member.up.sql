CREATE TABLE token_member (
    id SERIAL PRIMARY KEY,
    token_id INT REFERENCES token(id) ON DELETE CASCADE, 
    name VARCHAR(64),
    sort INT,
    description_text TEXT,
    description_image_file VARCHAR(255),--TODO image file
);
