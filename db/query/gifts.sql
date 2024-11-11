-- name: InsertGift :one
INSERT INTO gifts (id, gifter, recipient, message, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetGift :one
SELECT * FROM gifts WHERE id = $1 LIMIT 1;