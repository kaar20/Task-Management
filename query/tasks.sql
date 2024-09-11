-- name: CreateTask :exec
INSERT INTO tasks (title, description, status, priority, due_date, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetTaskByID :one
SELECT id, title, description, status, priority, due_date, user_id
FROM tasks
WHERE id = $1;
-- name: TasksList :many
SELECT * FROM tasks;
-- WHERE id = $1;


-- name: GetTasksByUserID :many
SELECT id, title, description, status, priority, due_date, user_id
FROM tasks
WHERE user_id = $1;

-- name: UpdateTask :exec
UPDATE tasks
SET title = $1, description = $2, status = $3, priority = $4, due_date = $5, user_id = $6
WHERE id = $7
RETURNING id;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;
