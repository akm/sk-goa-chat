-- +goose Up
CREATE TABLE `channels` (
    id serial PRIMARY KEY,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    name VARCHAR(255) NOT NULL,
    visibility ENUM('public', 'private') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE channels;
