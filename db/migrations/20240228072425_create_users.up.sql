create extension if not exists citext;

create table users (
  id bigint primary key generated by default as identity,
  email citext not null unique,
  password_hash varchar(255),
  inserted_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);
