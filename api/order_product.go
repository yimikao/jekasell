package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jekasell/db/sqlc"
)

type createOrderProductRequest struct {
	OrderID   int64 `json:"product_id" binding:"required,min=1"`
	ProductID int64 `json:"product_id" binding:"required,min=1"`
}

func (s *Server) CreateOrderProduct(ctx *gin.Context) {
	var req createOrderProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	op, err := s.store.CreateOrderProduct(ctx, db.CreateOrderProductParams{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, op)
}

type listOrderProductRequest struct {
	OrderID   int64 `json:"product_id" binding:"required,min=1"`
	ProductID int64 `json:"product_id" binding:"required,min=1"`
}

func (s *Server) ListOrderProducts(ctx *gin.Context) {
	var req listOrderProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	op, err := s.store.ListOrderProducts(ctx, db.ListOrderProductsParams{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, op)
}
