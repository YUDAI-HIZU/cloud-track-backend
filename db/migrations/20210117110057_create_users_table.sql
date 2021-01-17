
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `email` varchar(255) UNIQUE DEFAULT NULL,
    `encrypted_password` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;