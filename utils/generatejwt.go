package utils

import (
	"os"
	"time"

	"gin.com/aishwary11/models"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
