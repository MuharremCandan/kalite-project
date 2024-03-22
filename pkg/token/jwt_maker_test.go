package token

import (
	"go-backend-test/pkg/utils"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// 87.9%
func TestJWTMaker(t *testing.T) {
	keySize := 32
	maker_wrong, err := NewJWTMaker(utils.RandomString(1))
	require.Error(t, err)
	require.Nil(t, maker_wrong)

	maker, err := NewJWTMaker(utils.RandomString(keySize))
	require.NoError(t, err)

	username := utils.RandomUserName(6)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	userID, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, userID)

	token, err := maker.CreateToken(userID, username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.NoError(t, err)
	require.NotNil(t, payload)

	//check payload items
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.UserName)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Hour)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Hour)

}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	userID, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, userID)
	token, err := maker.CreateToken(userID, utils.RandomUserName(5), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	userID, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, userID)

	payload := NewPayload(userID, utils.RandomUserName(5), time.Minute)
	require.NoError(t, err)

	// testing the keyFunc
	token := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(tokenString)
	require.NoError(t, err)

	payload, err = maker.ValidateToken(tokenString)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
