CREATE TABLE user_info (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES user(id) ON DELETE CASCADE,
    investor_type VARCHAR(32) NOT NULL,
    auth_status VARCHAR(32) NOT NULL,
    user_jurisdiction VARCHAR(64),
    occupation VARCHAR(64),
    birthday DATE,
    residence VARCHAR(64),
    nationality VARCHAR(64),
    country_of_birth VARCHAR(64),
    authority_id INT REFERENCES authority(id) ON DELETE CASCADE,
);
