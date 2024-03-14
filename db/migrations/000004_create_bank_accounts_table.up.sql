CREATE TABLE bank_accounts (
    bank_account_id SERIAL PRIMARY KEY,
    bank_name VARCHAR(100),
    bank_account_name VARCHAR(100),
    bank_account_number VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);