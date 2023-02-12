package server

import (
	"github.com/OpenSauce/png/golang/store"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db *store.Store
}

func New(addr string, db *store.Store) *Server {
	router := gin.Default()
	initializeRoutes(router)
	router.Run(addr)
	return &Server{
		db: db,
	}
}

func initializeRoutes(router *gin.Engine) {
	router.GET("/", nil)
}
