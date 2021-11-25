-- name: CreateOrder :one
INSERT INTO orders (
    user_id
) VALUES (
    $1
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1;

-- name: ListOrders :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY id;

