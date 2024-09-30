CREATE TABLE admin_bank (
    id SERIAL PRIMARY KEY,
    bank_name VARCHAR(64) NOT NULL,
    legal_tender_id INT REFERENCES legal_tender(id) ON DELETE CASCADE,
    iban_code VARCHAR(34) NOT NULL,
    account_holder VARCHAR(64) NOT NULL,
    account_number VARCHAR(64) NOT NULL,
    bin VARCHAR(20) NOT NULL,-- business_identification_number
    bic VARCHAR(20) NOT NULL, --bank_identifier_code
    beneficiary VARCHAR(64),
    kbe VARCHAR(10)
);
