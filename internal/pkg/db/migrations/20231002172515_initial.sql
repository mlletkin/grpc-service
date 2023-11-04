-- Thank you for giving goose a try!
-- 
-- This file was automatically created running goose init. If you're familiar with goose
-- feel free to remove/rename this file, write some SQL and goose up. Briefly,
-- 
-- Documentation can be found here: https://pressly.github.io/goose
--
-- A single goose .sql file holds both Up and Down migrations.
-- 
-- All goose .sql files are expected to have a -- +goose Up annotation.
-- The -- +goose Down annotation is optional, but recommended, and must come after the Up annotation.
-- 
-- The -- +goose NO TRANSACTION annotation may be added to the top of the file to run statements 
-- outside a transaction. Both Up and Down migrations within this file will be run without a transaction.
-- 
-- More complex statements that have semicolons within them must be annotated with 
-- the -- +goose StatementBegin and -- +goose StatementEnd annotations to be properly recognized.
-- 
-- Use GitHub issues for reporting bugs and requesting features, enjoy!

-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
  id serial primary key not null,
  heading text not null,
  text text not null,
  likes_count int not null default 0
);
CREATE TABLE comments (
  id serial primary key not null,
  text text not null,
  likes_count int not null default 0,
  post_id int references posts (id) on delete cascade
);
-- +goose StatementEnd
-- +goose Down
DROP TABLE comments;
DROP TABLE posts;
