-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE roles (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       name VARCHAR(50) NOT NULL UNIQUE,
                       description TEXT
);

CREATE TABLE permissions (
                             id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             code VARCHAR(100) NOT NULL UNIQUE,
                             description TEXT
);

CREATE TABLE role_permissions (
                                  role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
                                  permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
                                  PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,
                       name VARCHAR(100),
                       role_id UUID REFERENCES roles(id) ON DELETE SET NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
