-- name: GetQuestion :one
SELECT * FROM question
WHERE id = $1 LIMIT 1;

-- name: ListQuestions :many
SELECT * FROM question
ORDER BY id;

-- name: CreateQuestion :one
INSERT INTO question (
    song
) VALUES (
    $1
)
RETURNING *;


-- name: GetChoice :one
SELECT * FROM choice
WHERE question_id = $1 LIMIT 4;

-- name: ListChoices :many
SELECT * FROM choice
ORDER BY id;

-- name: CreateChoice :one
INSERT INTO choice (
    question_id, choice, is_correct
) VALUES (
    $1, $2, $3
)
RETURNING *;


-- name: GetUserQuestionChoice :one
SELECT * FROM questionchoice
WHERE question_id = $1 LIMIT 1;

-- name: CreateUserQuestionChoice :one
INSERT INTO questionchoice (
    user_id, choice_id, question_id, is_correct
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: CreateGameSession :one
insert into gamesession (code, trivia_id, finished_at)
values ($1, $2, $3)
RETURNING *;

-- name: GetGameSession :one
select * from gamesession
where code = $1;

-- name: GetTrivia :one
select id, name from trivia
where name = $1;

-- name: GetAllTrivia :many
select id, name, image_path from trivia;

-- name: FetchTriviaByGameSession :many
SELECT
    "question"."song",
    "choice"."is_correct",
    "choice"."choice",
    "choice"."id"
FROM
    "trivia"
        INNER JOIN "gamesession" ON (
            "trivia"."id" = "gamesession"."trivia_id"
        )
        LEFT OUTER JOIN "trivia_questions" ON (
            "trivia"."id" = "trivia_questions"."trivia_id"
        )
        LEFT OUTER JOIN "question" ON (
            "trivia_questions"."question_id" = "question"."id"
        )
        LEFT OUTER JOIN "choice" ON (
            "question"."id" = "choice"."question_id"
        )
WHERE
        "gamesession"."id" = $1;