-- name: CreateUser :exec
INSERT INTO users (id, name) VALUES (?, ?);

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateProduct :exec
INSERT INTO products (id, name, user_id) VALUES (?, ?, ?);

-- name: GetProduct :one
SELECT * FROM products WHERE id = ?;

-- name: GetProducts :many
SELECT * FROM products;
