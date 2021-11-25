-- name: CreateOrderProduct :one
INSERT INTO order_products (
    order_id,
    product_id
) VALUES (
    $1, $2
) RETURNING *;


-- name: ListOrderProducts :many
SELECT * FROM order_products
WHERE
    order_id = $1 AND
    product_id = $2
ORDER BY id;

