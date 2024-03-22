package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrExpiredToken = fmt.Errorf("token has expired")
	ErrInvalidToken = fmt.Errorf("token is invalid")
)

// Payload is the payload data of a JWT token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"user_name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload a spesific username and duration
func NewPayload(id uuid.UUID, username string, duration time.Duration) *Payload {
	payload := &Payload{
		ID:        id,
		UserName:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload
}

// Valid checks if the token is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
