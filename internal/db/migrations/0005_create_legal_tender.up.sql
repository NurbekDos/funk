CREATE TABLE legal_tender (
    id SERIAL PRIMARY KEY,
    legal_tender_name VARCHAR(8) UNIQUE NOT NULL
);
-- Вставка фиксированных значений в таблицу legal_tender
INSERT INTO legal_tender (legal_tender_name) VALUES
('USD'),
('KZT'),
('RUB');