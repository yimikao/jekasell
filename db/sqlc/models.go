// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderProduct struct {
	OrderID   int64     `json:"order_id"`
	ProductID int64     `json:"product_id"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Quantity    int64          `json:"quantity"`
	Description string         `json:"description"`
	SellerID    int64          `json:"seller_id"`
	AvatarUrl   sql.NullString `json:"avatar_url"`
	CreatedAt   time.Time      `json:"created_at"`
}

type User struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	PhoneNumber sql.NullString `json:"phone_number"`
	Address     sql.NullString `json:"address"`
	CreatedAt   time.Time      `json:"created_at"`
}
