package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jekasell/db/sqlc"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := s.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, createdUser)

}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetUser(ctx *gin.Context) {
	var req getUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	u, err := s.store.GetUser(ctx, int64(req.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (s *Server) ListUsers(ctx *gin.Context) {
	us, err := s.store.ListUsers(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, us)
}

type updateUserRequest struct {
	ID       int64  `json"id"`
	Password string `json:"password"`
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	args := db.UpdateUserParams{
		ID:       req.ID,
		Password: req.Password,
	}
	u, err := s.store.UpdateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, u)
}
