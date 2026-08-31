package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	// 1. Get the cookie from the request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
		return
	}

	// 2. Decode and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return your secret key (ideally from environment variables)
		return []byte(os.Getenv("SECRET")), nil
	})

	// 3. Check for parsing errors or invalid tokens
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	// 4. Extract claims and verify expiration
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check expiration (golang-jwt v5 does this automatically in token.Valid, but here is a manual check if needed)
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Token expired"})
			return
		}

		// 5. Find user and attach the subject (sub) to the context
		// You can fetch full user data from your DB here if needed
		User := claims["sub"]
		c.Set("User", User)
	}

	// 6. Continue to the next handler
	c.Next()
}
