-- migrate:up
alter table trivia add column image_path text null;
-- migrate:down

