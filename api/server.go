package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/golang-projects/simplebank/db/sqlc"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store *db.Store
	router *gin.Engine
}


// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/",func(ctx *gin.Context) {ctx.JSON(http.StatusOK,gin.H{"message": "hello world"})})

	router.POST("/accounts",server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PUT("/accounts/:id", server.updateAccount)

	server.router = router 
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}


func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}