package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ithaquaKr/taskManager/db/sqlc"
)

// Server is a struct that holds the gin engine and the database store.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new server.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/tasks/:id", server.getTask)
	router.GET("/tasks", server.listTask)
	router.POST("/tasks", server.createTask)
	server.router = router
	return server
}

// Run starts the server.
func (server *Server) Run(address string) error {
	return server.router.Run(address)
}

// errorResponse returns a JSON response with the error message.
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
