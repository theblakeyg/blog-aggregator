-- +goose up
CREATE TABLE users (
    id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name VARCHAR(50) UNIQUE
);

-- +goose down
DROP TABLE users;