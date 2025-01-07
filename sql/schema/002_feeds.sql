-- +goose Up
CREATE TABLE feeds (
    id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name VARCHAR,
    url VARCHAR,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;