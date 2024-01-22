package config

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"Gotest/models"
)

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	return os.Getenv("JWTSECRET")
}

func GenerateToken(userInput models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userInput.Username,
		"email":    userInput.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}
