-- name: create-uuid-extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- name: create-products-table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL  PRIMARY KEY,
    guid UUID DEFAULT uuid_generate_v4(),
    name VARCHAR(255) UNIQUE NOT NULL,
    price REAL NOT NULL,
    description VARCHAR,
    created_at TEXT NOT NULL,
    updated_at TEXT
);