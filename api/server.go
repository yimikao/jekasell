package api

import (
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

	r.POST("/users", svr.CreateUser)
	r.GET("/users", svr.ListUsers)
	r.GET("/users/:id", svr.GetUser)
	svr.router = r
	return
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
