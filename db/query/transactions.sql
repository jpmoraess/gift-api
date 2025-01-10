-- name: InsertTransaction :one
INSERT INTO transactions (id, external_id, amount, date, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;