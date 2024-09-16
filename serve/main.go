package main

import (
	"context"
	"log"
	"net/http"
	"shared"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5" // Updated to jwt/v5
)

var keycloakClient *gocloak.GoCloak // Updated to pointer

var realm = "pcbb"                        // Replace with your Keycloak realm
var clientID = "your-backend-client"      // Replace with your Keycloak backend client ID
var keycloakURL = "http://localhost:8080" // Replace with Keycloak URL

func main() {
	// Initialize DB connection
	db := shared.ConnectDB()

	// Initialize Keycloak client
	keycloakClient = gocloak.NewClient(keycloakURL) // Now should work

	// Setup Gin router
	r := gin.Default()

	// Add token authentication middleware
	r.Use(TokenAuthMiddleware())

	// Define routes (example route)
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "This is a protected route",
		})
	})

	// Start the server
	r.Run(":8082")
}

// TokenAuthMiddleware validates the token in Authorization header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token with Keycloak
		isValid, claims, err := ValidateToken(token)
		if err != nil || !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Optionally set user information into the context for future use
		c.Set("user_id", claims["sub"]) // Keycloak's user ID (sub claim)
		c.Next()
	}
}

// ValidateToken checks if the token is valid using Keycloak
func ValidateToken(token string) (bool, jwt.MapClaims, error) {
	ctx := context.Background() // Create a context

	// Retrospect token with Keycloak
	rptResult, err := keycloakClient.RetrospectToken(ctx, token, clientID, "", realm)
	if err != nil {
		log.Printf("Failed to introspect token: %v", err)
		return false, nil, err
	}

	if !*rptResult.Active {
		return false, nil, nil
	}

	claims := jwt.MapClaims{}
	_, err = keycloakClient.DecodeAccessTokenCustomClaims(ctx, token, realm, &claims)
	if err != nil {
		return false, nil, err
	}

	return true, claims, nil
}
