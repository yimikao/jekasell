// Code generated by sqlc. DO NOT EDIT.
// source: order_product.sql

package db

import (
	"context"
)

const createOrderProduct = `-- name: CreateOrderProduct :one
INSERT INTO order_products (
    order_id,
    product_id
) VALUES (
    $1, $2
) RETURNING order_id, product_id, quantity, created_at
`

type CreateOrderProductParams struct {
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
}

func (q *Queries) CreateOrderProduct(ctx context.Context, arg CreateOrderProductParams) (OrderProduct, error) {
	row := q.db.QueryRowContext(ctx, createOrderProduct, arg.OrderID, arg.ProductID)
	var i OrderProduct
	err := row.Scan(
		&i.OrderID,
		&i.ProductID,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const listOrderProducts = `-- name: ListOrderProducts :many
SELECT order_id, product_id, quantity, created_at FROM order_products
WHERE
    order_id = $1 AND
    product_id = $2
ORDER BY id
`

type ListOrderProductsParams struct {
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
}

func (q *Queries) ListOrderProducts(ctx context.Context, arg ListOrderProductsParams) ([]OrderProduct, error) {
	rows, err := q.db.QueryContext(ctx, listOrderProducts, arg.OrderID, arg.ProductID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderProduct{}
	for rows.Next() {
		var i OrderProduct
		if err := rows.Scan(
			&i.OrderID,
			&i.ProductID,
			&i.Quantity,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}