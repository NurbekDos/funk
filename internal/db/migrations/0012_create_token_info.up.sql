CREATE TABLE token_info (
    id SERIAL PRIMARY KEY,
    token_id INT REFERENCES token(id) ON DELETE CASCADE, 
    title VARCHAR(64),
    sort INT,
    description_text TEXT,
    description_image_file VARCHAR(255),--TODO image file
    description_video_file VARCHAR(255),--TODO video file
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
