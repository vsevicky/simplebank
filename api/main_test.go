package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/vsevicky/simplebank/db/sqlc"
)

func newTestServer(store db.Store) *Server {
	server := NewServer(store)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
