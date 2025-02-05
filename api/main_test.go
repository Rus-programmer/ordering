package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "ordering/db/sqlc"
	"ordering/utils"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	server, err := NewServer(util.Config{}, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
