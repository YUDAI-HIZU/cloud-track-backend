
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users (
    id                 bigint(20) not null auto_increment primary key,
    email              varchar(255) unique default null,
    encrypted_password varchar(255) not null,
    name               varchar(255) not null,
    created_at         datetime not null,
    updated_at         datetime not null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;