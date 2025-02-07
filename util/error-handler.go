package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
)

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

var (
	ErrInvalidToken           = errors.New("token is invalid")
	ErrExpiredToken           = errors.New("token has expired")
	ErrInvalidPassword        = errors.New("invalid password")
	ErrIncorrectSessionUser   = errors.New("incorrect session user")
	ErrMismatchedSessionToken = errors.New("mismatched session token")
	ErrSessionExpired         = errors.New("expired session")
	ErrBlockedSession         = errors.New("blocked session")
	ErrRecordNotFound         = pgx.ErrNoRows
	ErrUniqueViolation        = &pgconn.PgError{
		Code: UniqueViolation,
	}
)

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}

func ErrorHandler(err error) (int, gin.H) {
	switch {
	case errors.Is(err, ErrInvalidPassword):
		return http.StatusUnauthorized, ErrorResponse(err)
	case errors.Is(err, ErrIncorrectSessionUser):
		return http.StatusUnauthorized, ErrorResponse(err)
	case errors.Is(err, ErrMismatchedSessionToken):
		return http.StatusUnauthorized, ErrorResponse(err)
	case errors.Is(err, ErrSessionExpired):
		return http.StatusUnauthorized, ErrorResponse(err)
	case errors.Is(err, ErrBlockedSession):
		return http.StatusUnauthorized, ErrorResponse(err)
	case errors.Is(err, ErrInvalidToken):
		return http.StatusUnauthorized, ErrorResponse(err)
	case errors.Is(err, ErrExpiredToken):
		return http.StatusUnauthorized, ErrorResponse(err)
	case errors.Is(err, ErrRecordNotFound):
		return http.StatusNotFound, ErrorResponse(err)
	case errors.Is(err, ErrRecordNotFound):
		return http.StatusNotFound, ErrorResponse(err)
	}

	errCode := ErrorCode(err)
	switch errCode {
	case UniqueViolation:
		return http.StatusConflict, ErrorResponse(err)
	case ForeignKeyViolation:
		return http.StatusForbidden, ErrorResponse(err)
	}

	return http.StatusInternalServerError, ErrorResponse(err)
}
