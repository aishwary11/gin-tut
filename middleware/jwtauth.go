package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gin.com/aishwary11/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ResponseHelper(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			utils.ResponseHelper(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			utils.ResponseHelper(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}
		c.Set("user", claims)
		c.Next()
	}
}
