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
func GenerateJWT(Email, Username, Role string) (string, string, error) {
	claims := &JWTClaim{
		Email:    Email,
		Username: Username,
		Role:     Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var jwtSecretKey = []byte(os.Getenv("API_SECRET"))

	token.SignedString(jwtSecretKey)
	tokenString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		panic(err)
	}

	claims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecretKey)

	return tokenString, refreshTokenString, nil
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

/*

-------------------------------------------- JWT login -------------------------------------------------------------------

1- We pass an email and we verify if it exist our db | if it doesn't exist we return a response to inform user

2- If it exist we continue and verify the password (the hashed version) if it correspond to the password of the user in the db

3- If yes, we generate a token and a refresh token, by passing data like email, username or what ever who belong to the user and we sign the token by a secret key

4- Now we implement a middleware which verify if we are connected or not

5- In all request to protected routes, we must use that middleware by verifying the if we have token in authorization and if that token is authentic by
validating it.

6-

What do we do if the token expire ?

The token has a short life time validity. We have to create a refresh token to prevent the case where it expires.

But what is a refresh token?

A refresh token is a token generated in the same time like the access token. When the access token expires, we use the refresh token to ask for a new access token

Therefore, the refresh must be signed by a secret key and we should validate it

If




*/
