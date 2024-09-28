package middlewares

import (
	"fmt"
	"goauth/initializers"
	"goauth/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// CheckAuth is a middleware function that authenticates the user based on the JWT token in the Authorization header
func CheckAuth(c *gin.Context) {
	// Extract the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Split the header into "Bearer" and the token
	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Extract the token string
	tokenString := authToken[1]

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used for signing
		return []byte(os.Getenv("SECRET")), nil
	})

	// Check for parsing errors
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Extract and validate claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check if the token has expired
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Fetch the user from the database
	var user models.User
	result := initializers.DB.Where("id = ?", claims["ID"]).First(&user)
	if result.Error != nil || user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Set the user in the context for later use
	c.Set("user", user)

	// Set SameSite attribute for enhanced security
	c.SetSameSite(http.SameSiteLaxMode)

	// Continue to the next middleware or handler
	c.Next()
}
