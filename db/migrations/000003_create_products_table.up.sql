CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL CHECK (LENGTH(name) >= 5 AND LENGTH(name) <= 60),
    price DECIMAL NOT NULL CHECK (price >= 0),
    image_url VARCHAR(255) NOT NULL CHECK (image_url ~* '^https?://.*$'),
    stock INT NOT NULL CHECK (stock >= 0),
    condition VARCHAR(10) NOT NULL CHECK (condition IN ('new', 'second')),
    tags TEXT[] NOT NULL DEFAULT '{}',
    is_purchaseable BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);