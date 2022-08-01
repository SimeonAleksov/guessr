-- migrate:up

CREATE TABLE "question" (
    "id" serial NOT NULL PRIMARY KEY,
    "song" varchar(100) NOT NULL,
    "is_active" bool NOT NULL,
    "created_at" date NOT NULL default now(),
    "updated_at" date NULL);

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
    "id" serial NOT NULL PRIMARY KEY ,
    "trivia_id" bigint NOT NULL REFERENCES "trivia" ("id") DEFERRABLE INITIALLY DEFERRED,
    "question_id" bigint NOT NULL REFERENCES "question" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TABLE "gamesession" (
    "id" serial NOT NULL PRIMARY KEY ,
    "code" varchar(128) NOT NULL,
    "trivia_id" bigint NOT NULL REFERENCES "trivia" ("id") DEFERRABLE INITIALLY DEFERRED,
    "created_at" date NOT NULL DEFAULT now(),
    "finished_at" date NULL
);

CREATE TABLE "questionchoice" (
    "id" serial NOT NULL PRIMARY KEY ,
    "is_correct" bool NOT NULL,
    "answer_time" date NOT NULL DEFAULT now(),
    "choice_id" bigint NOT NULL REFERENCES "choice" ("id") DEFERRABLE INITIALLY DEFERRED,
    "question_id" bigint NOT NULL REFERENCES "question" ("id") DEFERRABLE INITIALLY DEFERRED,
    "user_id" integer NOT NULL REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "gamesessionscoreboard" (
    "id" serial NOT NULL PRIMARY KEY,
    "score" integer NOT NULL,
    "game_session_id" bigint NOT NULL REFERENCES "gamesession" ("id") DEFERRABLE INITIALLY DEFERRED,
    "user_id" integer NOT NULL REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE UNIQUE INDEX "trivia_questions_trivia_id_question_id_7d8bc5ac_uniq" ON "trivia_questions" ("trivia_id", "question_id");
CREATE INDEX "trivia_questions_trivia_id_9ffcbec1" ON "trivia_questions" ("trivia_id");
CREATE INDEX "trivia_questions_question_id_43ce0a83" ON "trivia_questions" ("question_id");
CREATE INDEX "questionchoice_choice_id_20d69ef0" ON "questionchoice" ("choice_id");
CREATE INDEX "questionchoice_question_id_6b81dda8" ON "questionchoice" ("question_id");
CREATE INDEX "questionchoice_user_id_01de416e" ON "questionchoice" ("user_id");
CREATE INDEX "gamesessionscoreboard_game_session_id_eeb5ceac" ON "gamesessionscoreboard" ("game_session_id");
CREATE INDEX "gamesessionscoreboard_user_id_6af55026" ON "gamesessionscoreboard" ("user_id");
CREATE INDEX "gamesession_trivia_id_d30fb4fc" ON "gamesession" ("trivia_id");
CREATE INDEX "choice_question_id_f16e4a48" ON "choice" ("question_id");