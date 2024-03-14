CREATE TABLE payment (
    payment_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(product_id),
    bank_account_id INT NOT NULL REFERENCES bank_accounts(bank_account_id),
    payment_proof_image_url VARCHAR(255) NOT NULL CHECK (payment_proof_image_url ~* '^https?://.*$'),
    quantity INT NOT NULL CHECK (quantity >= 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);