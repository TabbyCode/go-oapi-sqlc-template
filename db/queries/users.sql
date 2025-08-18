-- name: GetUser :one
SELECT
    id,
    name,
    email,
    created_at,
    updated_at
FROM users
WHERE id = sqlc.arg('id')
LIMIT 1;

-- name: ListUsers :many
SELECT
    id,
    name,
    email,
    created_at,
    updated_at
FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (
    name,
    email
)
VALUES (sqlc.arg('name'), sqlc.arg('email'))
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    name = COALESCE(sqlc.narg('name'), name),
    email = COALESCE(sqlc.narg('email'), email),
    updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = sqlc.arg('id')
RETURNING *;
