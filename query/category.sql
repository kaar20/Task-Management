-- internal/db/queries/categories.sql

-- name: CreateCategory :exec
INSERT INTO categories (name)
VALUES ($1)
RETURNING id;

-- name: GetCategoryByID :one
SELECT id, name
FROM categories
WHERE id = $1;

-- name: UpdateCategory :exec
UPDATE categories
SET name = $1
WHERE id = $2
RETURNING id;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: ListCategories :many
SELECT id, name
FROM categories;
