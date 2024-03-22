package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const minSecretKeySize = 32

// JWTMaker is JSON Web Token Maker
type JWTMaker struct {
	secretkey string
}

func NewJWTMaker(secretkey string) (Maker, error) {
	if len(secretkey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretkey: secretkey}, nil
}

// CreateToken creates a new token for a spesific username and duration
func (maker *JWTMaker) CreateToken(id uuid.UUID, username string, duration time.Duration) (string, error) {
	payload := NewPayload(id, username, duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(maker.secretkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates a token is valid or not
func (maker *JWTMaker) ValidateToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretkey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
