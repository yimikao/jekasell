package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jekasell/db/sqlc"
)

type createUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone"`
	Address     string `json:"address"`
}

type createUserResponse struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewUserResponse(user db.User) createUserResponse {
	return createUserResponse{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
		Address:     user.Address.String,
		CreatedAt:   user.CreatedAt,
	}
}
func (s *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.CreateUserParams{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: sql.NullString{String: req.PhoneNumber, Valid: true},
		Address:     sql.NullString{String: req.Address, Valid: true},
	}

	createdUser, err := s.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := NewUserResponse(createdUser)
	ctx.JSON(http.StatusCreated, rsp)

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
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (s *Server) ListUsers(ctx *gin.Context) {
	us, err := s.store.ListUsers(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, us)
}

type updateUserRequest struct {
	ID       int64  `json:"id"`
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
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (s *Server) DeleteUser(ctx *gin.Context) {

	var req getUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := s.store.GetUser(ctx, int64(req.ID))
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = s.store.DeleteUser(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"messages": "User deleted successfully!"})
}

type loginUserRequest struct {
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type loginUserResponse struct {
	AccessToken string             `json:"access_token"`
	User        createUserResponse `json:"user"`
}

func (s *Server) LoginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	u, err := s.store.GetUserByEmail(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if u.Password != req.Password {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	accessToken, err := s.tokenMaker.CreateToken(u.Name, s.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        NewUserResponse(u),
	}
	ctx.JSON(http.StatusOK, rsp)
}
