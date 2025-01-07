-- +goose Up
CREATE TABLE users (
    id UUID UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name VARCHAR(50) UNIQUE
);

-- +goose Down
DROP TABLE users;