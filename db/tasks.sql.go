// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tasks.sql

package db

import (
	"context"
	"database/sql"
)

const createTask = `-- name: CreateTask :exec
INSERT INTO tasks (title, description, status, priority, due_date, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`

type CreateTaskParams struct {
	Title       string
	Description sql.NullString
	Status      sql.NullString
	Priority    sql.NullString
	DueDate     sql.NullTime
	UserID      sql.NullInt32
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) error {
	_, err := q.db.ExecContext(ctx, createTask,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.Priority,
		arg.DueDate,
		arg.UserID,
	)
	return err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, title, description, status, priority, due_date, user_id
FROM tasks
WHERE id = $1
`

func (q *Queries) GetTaskByID(ctx context.Context, id int32) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.DueDate,
		&i.UserID,
	)
	return i, err
}

const getTasksByUserID = `-- name: GetTasksByUserID :many


SELECT id, title, description, status, priority, due_date, user_id
FROM tasks
WHERE user_id = $1
`

// WHERE id = $1;
func (q *Queries) GetTasksByUserID(ctx context.Context, userID sql.NullInt32) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasksByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Priority,
			&i.DueDate,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const tasksList = `-- name: TasksList :many
SELECT id, title, description, status, priority, due_date, user_id FROM tasks
`

func (q *Queries) TasksList(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, tasksList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Priority,
			&i.DueDate,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :exec
UPDATE tasks
SET title = $1, description = $2, status = $3, priority = $4, due_date = $5, user_id = $6
WHERE id = $7
RETURNING id
`

type UpdateTaskParams struct {
	Title       string
	Description sql.NullString
	Status      sql.NullString
	Priority    sql.NullString
	DueDate     sql.NullTime
	UserID      sql.NullInt32
	ID          int32
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateTask,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.Priority,
		arg.DueDate,
		arg.UserID,
		arg.ID,
	)
	return err
}
