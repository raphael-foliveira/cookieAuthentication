package authorization

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/cookieAuthentication/api/models"
)

type CustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("super-secret-key")

func GenerateToken(user models.User) (string, error) {
	claims := CustomClaims{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "authorization-server",
			Subject:   user.Username,
			ID:        user.Id,
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newToken.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	var customClaims CustomClaims

	parsedToken, err := jwt.ParseWithClaims(tokenString, &customClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, err
	}
	return &customClaims, nil
}
