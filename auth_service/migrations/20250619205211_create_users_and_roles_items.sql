-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE roles (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,
                       name VARCHAR(100),
                       role_id INTEGER,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
