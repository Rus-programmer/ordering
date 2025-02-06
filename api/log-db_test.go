package api

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/logging"
	util "ordering/utils"
	"testing"
	"time"
)

func TestLogDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockdb.NewMockStore(ctrl)

	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, mockStore)
	require.NoError(t, err)

	server.router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	ginInfo := &logging.LogInfo{
		Method:     "GET",
		Path:       "/test",
		StatusCode: http.StatusOK,
		Duration:   time.Millisecond * 50,
		BeginTime:  time.Now(),
	}

	done := make(chan struct{})

	mockStore.EXPECT().
		CreateLog(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, arg db.CreateLogParams) error {
			require.Equal(t, ginInfo.Method, arg.Method)
			require.Equal(t, ginInfo.Path, arg.Path)
			require.Equal(t, int32(ginInfo.StatusCode), arg.StatusCode)
			elapsedTime, err := time.ParseDuration(arg.ElapsedTime)
			require.NoError(t, err)
			require.InDelta(t, ginInfo.Duration.Milliseconds(), elapsedTime.Milliseconds(), 1000, "Duration mismatch")
			close(done)
			return nil
		}).Times(1)

	req, _ := http.NewRequest(http.MethodGet, "/test", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	server.router.ServeHTTP(w, req)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Timeout waiting for CreateLog call")
	}

	require.Equal(t, http.StatusOK, w.Code)
}
