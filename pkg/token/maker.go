package token

import (
	"time"

	"github.com/google/uuid"
)

// Maker is an interface that managing tokens
type Maker interface {
	//CreateToken creates a new token for a spesific username and duration
	CreateToken(id uuid.UUID, username string, duration time.Duration) (string, error)

	//ValidateToken validates a token is valid or not
	ValidateToken(token string) (*Payload, error)
}
