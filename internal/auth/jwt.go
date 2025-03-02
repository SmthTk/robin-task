package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWTService is a service to generate and validate JWT tokens
type JWTService struct {
	SecretKey string
	Expiry    int
}

// NewJWTService creates a new JWTService with a secret key and expiry time
func NewJWTService(secretKey string, expiry int) *JWTService {
	return &JWTService{
		SecretKey: secretKey,
		Expiry:    expiry,
	}
}

// GenerateToken generates a new JWT token with userID and role as claims
func (s *JWTService) GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(time.Duration(s.Expiry) * time.Hour).Unix(), // expiry in hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates the token and returns userID and role if valid
func (s *JWTService) ValidateToken(tokenString string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("could not parse claims")
	}

	userID := uint(claims["userID"].(float64)) // Convert from float64 to uint
	role := claims["role"].(string)

	return userID, role, nil
}
