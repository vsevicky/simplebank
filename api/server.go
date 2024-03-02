package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vsevicky/simplebank/db/sqlc"
)

// Server servers HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer Creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	// Add routes to the router
	server.router = router
	return server
}

// Start runs the HTTP server on the specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponses(err error) gin.H {
	return gin.H{"error": err.Error()}
}
