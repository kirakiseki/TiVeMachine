package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"user/setup"
)

type JWTClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

var authPaths = []string{
	"/api/user/info",
	"/api/user/setDescription",
	"/api/user/setSex",
	"/api/user/setAvatar",
}

func AuthInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAuthPath := false

		for _, path := range authPaths {
			if c.FullPath() == path {
				isAuthPath = true
				break
			}
		}

		if !isAuthPath {
			c.Next()
			return
		}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status_code": 50000, "message": "Unauthorized"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(setup.Config.JWTSecret), nil
		}, jwt.WithLeeway(5*time.Second))

		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			c.Set("userID", claims.ID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status_code": 50001, "message": err.Error()})
			return
		}
	}
}
