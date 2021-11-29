package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jekasell/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(s db.Store) (svr *Server) {
	svr = &Server{store: s}
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to jekasell"})
	})
	r.GET("/auth/signup")
	r.GET("/auth/signin")

	r.POST("/users", svr.CreateUser)
	r.GET("/users", svr.ListUsers)
	r.PUT("/users", svr.UpdateUser)
	r.GET("/users/:id", svr.GetUser)
	r.DELETE("/users/:id", svr.DeleteUser)

	r.POST("/products", svr.CreateProduct)
	r.GET("/products", svr.ListProducts)
	r.PUT("/products", svr.UpdateProduct)
	r.GET("/products/:id", svr.GetProduct)
	r.DELETE("/products/:id", svr.DeleteProduct)

	r.POST("/orders", svr.CreateOrder)
	r.GET("/orders", svr.ListOrders)
	r.GET("/orders/:id", svr.GetOrder)

	r.POST("/orderproducts", svr.CreateOrderProduct)
	r.GET("/orderproducts", svr.ListOrderProducts)
	svr.router = r
	return
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
