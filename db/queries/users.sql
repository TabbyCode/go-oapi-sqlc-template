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
WHERE
    (sqlc.narg('name')::text IS NULL OR name = sqlc.narg('name')::text)
    AND (sqlc.narg('email')::text IS NULL OR email = sqlc.narg('email')::text)
ORDER BY created_at
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

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

-- name: DeleteUser :execrows
DELETE
FROM users
WHERE id = sqlc.arg('id');
