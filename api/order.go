package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createOrderRequest struct {
	UserID int64 `json:"id" binging:"required,min=1"`
}

func (s *Server) CreateOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	o, err := s.store.CreateOrder(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, o)
}

type getOrderRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetOrder(ctx *gin.Context) {
	var req getOrderRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	o, err := s.store.GetOrder(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
	}
	ctx.JSON(http.StatusOK, o)
}

func (s *Server) ListOrders(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	os, err := s.store.ListOrders(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, os)
}
