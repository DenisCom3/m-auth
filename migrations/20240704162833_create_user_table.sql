-- +goose Up
create table IF NOT EXISTS users (
                       id serial primary key,
                       name varchar(50) not null,
                       email varchar(50) not null,
                       password varchar(200) not null,
                       role integer not null,
                       created_at timestamp not null default now(),
                       updated_at timestamp
);

-- +goose Down
drop table users;
