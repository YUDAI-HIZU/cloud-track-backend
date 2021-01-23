
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table profiles (
  user_id          bigint(20) not null primary key,
  introduction     varchar(255) default null,
  web_url          varchar(255) default null,
  icon_image_name  varchar(255) unique default null,
  cover_image_name varchar(255) unique default null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table profiles;