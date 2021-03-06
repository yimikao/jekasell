-- name: CreateProduct :one
INSERT into products (
    name,
    quantity,
    description,
    seller_id,
    avatar_url
) VALUES (
    $1, $2, $3, $4, $5
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