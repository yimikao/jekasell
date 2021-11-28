package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jekasell/db/sqlc"
)

type getProductRequest struct {
	ID int64 `uri:"id" binding:'required,min=1"`
}

func (s *Server) GetProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	p, err := s.store.GetProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, p)

}

func (s *Server) ListProducts(ctx *gin.Context) {
	ps, err := s.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, ps)
}

type createProductRequest struct {
	Name        string `json:"name"`
	Quantity    int64  `json:"quantity"`
	Description string `json:"description"`
	SellerID    int64  `json:"seller_id"`
}

func (s *Server) CreateProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	cp, err := s.store.CreateProduct(ctx, db.CreateProductParams{
		Name:        req.Name,
		Quantity:    req.Quantity,
		Description: req.Description,
		SellerID:    req.SellerID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, cp)
}

type updateProductRequest struct {
	ID       int64 `json:"id" binding:"required,min=1"`
	Quantity int64 `json:"quantity" binding:"required,min=1"`
}

func (s *Server) UpdateProduct(ctx *gin.Context) {
	var req updateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	up, err := s.store.UpdateProduct(ctx, db.UpdateProductParams{
		ID:       req.ID,
		Quantity: req.Quantity,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, up)
}

func (s *Server) DeleteProduct(ctx *gin.Context) {
	var req getProductRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := s.store.GetProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNoContent, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = s.store.DeleteProduct(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "product deleted successfully!"})
}
