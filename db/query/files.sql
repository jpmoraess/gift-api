-- name: InsertFile :one
INSERT INTO files (id, name, extension, size, path)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFile :one
SELECT * FROM files WHERE id = $1 LIMIT 1;

-- name: DeleteFile :exec
DELETE FROM files WHERE id = $1;