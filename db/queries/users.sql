-- name: GetUser :one
SELECT
    id,
    name,
    email,
    created_at,
    updated_at
FROM users
WHERE id = $1
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
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    name = COALESCE($2, name),
    email = COALESCE($3, email),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;
