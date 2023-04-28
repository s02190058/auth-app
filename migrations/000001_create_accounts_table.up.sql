CREATE TABLE IF NOT EXISTS accounts
(
    id                 BIGSERIAL PRIMARY KEY,
    username           TEXT      NOT NULL UNIQUE,
    email              TEXT      NOT NULL UNIQUE,
    encrypted_password TEXT      NOT NULL,
    is_verified        BOOLEAN   NOT NULL,
    created_at         TIMESTAMP NOT NULL,
    updated_at         TIMESTAMP NOT NULL
);