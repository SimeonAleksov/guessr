CREATE TABLE "user" (
    id       BIGSERIAL PRIMARY KEY,
    username text      NOT NULL,
    password text      NOT NULL,
    created_at  date NOT NULL default now()
);