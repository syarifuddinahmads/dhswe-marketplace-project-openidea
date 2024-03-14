CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(15) NOT NULL CHECK (LENGTH(username) >= 5 AND LENGTH(username) <= 15),
    name VARCHAR(50) NOT NULL CHECK (LENGTH(name) >= 5 AND LENGTH(name) <= 50),
    password VARCHAR(15) NOT NULL CHECK (LENGTH(password) >= 5 AND LENGTH(password) <= 15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);