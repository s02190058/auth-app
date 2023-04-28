CREATE TABLE IF NOT EXISTS refresh_tokens
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT    NOT NULL REFERENCES accounts (id) ON DELETE CASCADE,
    token      TEXT      NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL
);

CREATE INDEX ON refresh_tokens (user_id);

CREATE INDEX ON refresh_tokens (token);