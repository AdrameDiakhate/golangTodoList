package utils

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type AuthInfos struct {
	Email    string
	Password string
}
type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"roleNom"`
	jwt.StandardClaims
}

//Generate a token
func GenerateJWT(Email, Username, Role string) (string, error) {
	claims := &JWTClaim{
		Email:    Email,
		Username: Username,
		Role:     Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var jwtSecretKey = []byte("JWT_SECRET_KEY")

	tokenString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		return "", err
	}

	token.SignedString([]byte(os.Getenv("API_SECRET")))
	return tokenString, nil
}

//Used to verify if the token provided is valid
func ValidateToken(signedToken string) (Email string, Username string, Role string, err error) {
	var jwtSecretKey = []byte("JWT_SECRET_KEY")

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
	)
	if err != nil {
		return "", "", "", err
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("couldn't parse claims")
		return "", "", "", err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", "", "", err
	}

	return claims.Email, claims.Username, claims.Role, err
}
