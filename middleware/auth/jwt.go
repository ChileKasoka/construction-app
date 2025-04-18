package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// jwtSecret is loaded in init()
var jwtSecret []byte

func init() {
	// Load .env file if needed (optional but common in local development)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading it (optional)")
	}

	jwts := os.Getenv("JWT_SECRET")
	if jwts == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	jwtSecret = []byte(jwts)
}

// CreateJWT creates a JWT token with custom claims
func CreateJWT(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseJWT validates a token string and returns claims
func ParseJWT(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}

	return claims, nil
}
