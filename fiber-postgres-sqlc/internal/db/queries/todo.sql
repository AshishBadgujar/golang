-- name: CreateTodo :one
INSERT INTO todos (title, completed)
VALUES ($1, $2)
RETURNING *;

-- name: GetTodos :many
SELECT * FROM todos
ORDER BY created_at DESC;

-- name: GetTodoByID :one
SELECT * FROM todos
WHERE id = $1;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2,
    completed = $3,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;


-- Each -- name: becomes a typed Go method.