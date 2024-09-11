-- name: AddTaskCategory :exec
INSERT INTO task_categories (task_id, category_id)
VALUES ($1, $2);

-- name: RemoveTaskCategory :exec
DELETE FROM task_categories
WHERE task_id = $1 AND category_id = $2;

-- name: GetCategoriesForTask :many
SELECT c.id, c.name
FROM categories c
JOIN task_categories tc ON c.id = tc.category_id
WHERE tc.task_id = $1;

-- name: GetTasksForCategory :many
SELECT t.id, t.title, t.description, t.status, t.priority, t.due_date, t.user_id
FROM tasks t
JOIN task_categories tc ON t.id = tc.task_id
WHERE tc.category_id = $1;
