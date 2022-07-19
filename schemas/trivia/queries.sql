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