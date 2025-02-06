package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID         uuid.UUID `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Role       string    `json:"role"`
	IssuedAt   time.Time `json:"issued_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific customerID and duration
func NewPayload(customerID int64, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:         tokenID,
		CustomerID: customerID,
		Role:       role,
		IssuedAt:   time.Now(),
		ExpiredAt:  time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
