CREATE TABLE token_receival_timing (
    id SERIAL PRIMARY KEY,
    token_receival_timing_name VARCHAR(64) UNIQUE NOT NULL
);

INSERT INTO token_receival_timing (token_receival_timing_name) VALUES
('At the time of purchase'),
('When application period ends');