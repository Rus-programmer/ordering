package api

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	"ordering/middleware"
	mockService "ordering/services/mock"
	"ordering/token"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"ordering/util"
)

func newTestServer(t *testing.T) (*Server, *mockService.MockService) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
		Environment:         "test",
	}
	store := mockdb.NewMockStore(ctrl)
	mockTokenMaker, err := token.NewPasetoMaker("12345678912345678912345678900032")
	require.NoError(t, err)
	newMiddleware := middleware.NewMiddleware(store, mockTokenMaker)

	service := mockService.NewMockService(ctrl)
	service.EXPECT().GetTokenMaker().Return(mockTokenMaker).AnyTimes()

	server, err := NewServer(config, newMiddleware, service)

	require.NoError(t, err)
	return server, service
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
