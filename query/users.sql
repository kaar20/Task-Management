-- name: CreateUser :exec
INSERT INTO users (username, password_hash, role)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetUserByID :one
SELECT id, username, role
FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET username = $1, password_hash = $2, role = $3
WHERE id = $4
RETURNING id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
