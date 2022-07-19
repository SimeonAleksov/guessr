-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM "user"
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY username;

-- name: CreateUser :one
INSERT INTO "user" (
    username, password
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
