package middleware

import (
	"lore_project/pkg/auth"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function to authenticate requests.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Xác thực token JWT
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := auth.VerifyJWT(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
            return
        }
		if claims["role"] != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
			c.Abort()
			return
		}
		
        c.Next()

	}
}

// LoggerMiddleware is a middleware function to log requests.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform logging logic here
		// For example, log the request method, URL, and other details

		// Continue to the next handler
		c.Next()
	}
}

