-- name: CreateProduct :one
INSERT into products (
    name,
    quantity,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: UpdateProduct :one
UPDATE products
SET quantity = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;