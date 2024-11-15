-- name: CreateUser :one
INSERT INTO users (id, username, password, full_name, email, created_at)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;