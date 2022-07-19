CREATE TABLE "question" (
    "id" serial NOT NULL PRIMARY KEY,
    "song" varchar(100) NOT NULL,
    "is_active" bool NOT NULL,
    "created_at" date NOT NULL default now(),
    "updated_at" date NULL
);

CREATE TABLE "choice" (
    "id" serial NOT NULL PRIMARY KEY,
    "is_correct" bool NOT NULL,
    "choice" varchar(256) NOT NULL,
    "question_id" bigint NOT NULL REFERENCES "question" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "trivia" (
    "id" serial NOT NULL PRIMARY KEY ,
    "name" varchar(128) NOT NULL,
    "created_at" date NOT NULL,
    "updated_at" date NULL
);

CREATE TABLE "trivia_questions" (
    "id" serial NOT NULL PRIMARY KEY,
    "trivia_id" bigint NOT NULL REFERENCES "trivia" ("id") DEFERRABLE INITIALLY DEFERRED,
    "question_id" bigint NOT NULL REFERENCES "question" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "gamesession" (
    "id" serial NOT NULL PRIMARY KEY,
    "code" varchar(128) NOT NULL,
    "trivia_id" bigint NOT NULL REFERENCES "trivia" ("id") DEFERRABLE INITIALLY DEFERRED,
    "created_at" date NOT NULL DEFAULT now(),
    "finished_at" date NOT NULL
);

CREATE TABLE "questionchoice" (
    "id" serial NOT NULL PRIMARY KEY,
    "is_correct" bool NOT NULL,
    "answer_time" date NOT NULL DEFAULT now(),
    "choice_id" bigint NOT NULL REFERENCES "choice" ("id") DEFERRABLE INITIALLY DEFERRED,
    "question_id" bigint NOT NULL REFERENCES "question" ("id") DEFERRABLE INITIALLY DEFERRED,
    "user_id" integer NOT NULL REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "gamesessionscoreboard" (
    "id" serial NOT NULL PRIMARY KEY,
    "score" integer NOT NULL,
    "game_session_id" bigint NOT NULL REFERENCES "gamesession" ("id") DEFERRABLE INITIALLY DEFERRED,
    "user_id" integer NOT NULL REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED
);
