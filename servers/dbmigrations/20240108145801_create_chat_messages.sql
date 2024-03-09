-- +goose Up
CREATE TABLE `chat_messages` (
    id serial PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    channel_id bigint(20) UNSIGNED NOT NULL,
    user_id bigint(20) UNSIGNED NULL,
    user_name VARCHAR(255) NOT NULL,
    content MEDIUMTEXT NOT NULL,
    CONSTRAINT fk_chat_messages_channel_id FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    CONSTRAINT fk_chat_messages_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE chat_messages;
