package logic

import (
	"bookManagement/db"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// exp in hours
type JwtManager struct {
	exp          time.Duration
	jwtSecretKey []byte
}

type claims struct {
	jwt.RegisteredClaims
	UserName string `json:"userName"`
}

func generateRandomKey() ([]byte, error) {
	// jwtKey := make([]byte, 32)
	jwtKey := []byte{74, 230, 12, 205, 34, 107, 180, 161, 170, 121, 8, 191, 71, 155, 180, 39, 169, 196, 85, 135, 196, 68, 232, 31, 10, 2, 190, 222, 47, 89, 144, 54}
	// if _, err := rand.Read(jwtKey); err != nil {
	// 	return nil, err
	// }
	return jwtKey, nil
}

func NewJwtManager(exp time.Duration) (*JwtManager, error) {
	secretKey, err := generateRandomKey()
	if err != nil {
		return nil, err
	}

	return &JwtManager{
		exp:          exp,
		jwtSecretKey: secretKey,
	}, nil
}

func (jwtManager *JwtManager) createNewToken(userName string) (string, error) {
	claims := &claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * jwtManager.exp)),
		},
		UserName: userName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtManager.jwtSecretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (jwtManager *JwtManager) verifyToken(tokenString string, dbStruct *db.Db) (string, error) {
	claims := &claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtManager.jwtSecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrTokenExpired {
			return "", errors.New("token has been expired")
		}
		return "", errors.New("invalid Token")
	}

	_, err = dbStruct.FindUserByUserName(claims.UserName)
	if err != nil {
		return "", errors.New("user not found")
	}

	return claims.UserName, nil
}
