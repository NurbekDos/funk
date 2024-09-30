CREATE TABLE authority (
    id SERIAL PRIMARY KEY,
    authority_name VARCHAR(32) UNIQUE NOT NULL
);

-- Вставка фиксированных значений в таблицу authority
INSERT INTO authority (authority_name) VALUES
('Issuer'),
('KYC Authorizer'),
('Asset Manager'),
('Case Manager'),
('Auditor');
