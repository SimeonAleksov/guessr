-- name: GetQuestion :one
SELECT * FROM questions
WHERE id = $1 LIMIT 1;

-- name: ListQuestions :many
SELECT * FROM questions
ORDER BY id;

-- name: CreateQuestion :one
INSERT INTO questions (
    question
) VALUES (
             $1
         )
RETURNING *;


-- name: GetChoice :one
SELECT * FROM choices
WHERE question_id = $1 LIMIT 4;

-- name: ListChoices :many
SELECT * FROM choices
ORDER BY id;

-- name: CreateChoice :one
INSERT INTO choices (
    question_id, choice, is_correct
) VALUES (
             $1, $2, $3
         )
RETURNING *;


-- name: GetUserQuestionChoice :one
SELECT * FROM user_question_choice
WHERE question_id = $1 LIMIT 1;

-- name: CreateUserQuestionChoice :one
INSERT INTO user_question_choice (
    user_id, choice_id, question_id, is_correct
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;
