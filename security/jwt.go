package security

import (
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Token represents a parsed and validated JWT token
type Token struct {
	UserID string
	Token  string
}

// JWTOptions options to create a JWTManager
type JWTOptions struct {
	SigningMethod  string
	PublicKeyPath  string
	PrivateKeyPath string
	Expiration     time.Duration
}

// JWTManager is the manager of JWT autentication
type JWTManager struct {
	privateKey []byte
	publicKey  []byte
	Options    JWTOptions
}

// CreateJWTManager create a JWTManager with the given options
func CreateJWTManager(o JWTOptions) *JWTManager {
	return &JWTManager{
		privateKey: []byte(o.PrivateKeyPath),
		publicKey:  []byte(o.PublicKeyPath),
		Options:    o,
	}
}

// GenerateToken Generates a JSON Web Token given an userId
func (m *JWTManager) GenerateToken(userID string) (string, error) {
	now := time.Now()
	claims := jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(m.Options.Expiration).Unix(),
		Subject:   userID,
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod(m.Options.SigningMethod), claims)
	tokenString, err := t.SignedString(m.privateKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenString, nil
}

// ValidateToken verify if the token given is a valid one
func (m *JWTManager) ValidateToken(tokenString string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		switch token.Method {
		case jwt.SigningMethodHS256:
			return m.privateKey, nil
		case jwt.SigningMethodRS256:
			return m.publicKey, nil
		default:
			return nil, errors.New("JWT Token is not Valid")
		}
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				return nil, errors.New("Token Expired, get a new one")
			default:
				return nil, errors.New("JWT Token ValidationError")
			}
		}
		return nil, errors.New("JWT Token Error Parsing the token or empty token")
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("JWT Token is not Valid")
	}

	userID := claims.Subject
	return &Token{UserID: userID, Token: tokenString}, nil
}
