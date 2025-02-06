package api

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "ordering/db/sqlc"
	"ordering/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
		Environment:         "test",
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
