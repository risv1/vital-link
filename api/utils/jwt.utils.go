package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"github.com/joho/godotenv"
)

var jwtSecret string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	jwtSecret = os.Getenv("JWT_SECRET")
}

func GenerateJWT(id string) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}