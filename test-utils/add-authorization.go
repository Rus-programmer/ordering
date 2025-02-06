package test_utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	db "ordering/db/sqlc"
	"ordering/token"
	"testing"
	"time"
)

func AddAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	customerID int64,
	role db.UserRole,
	duration time.Duration,
) {
	createdToken, payload, err := tokenMaker.CreateToken(customerID, string(role), duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, createdToken)
	request.Header.Set("authorization", authorizationHeader)
}
