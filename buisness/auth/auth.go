package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vishn007/go-service-template/foundation/logger"
)

var jwtSecret string = "3242342adsfadqwerdvsdfaerwrw3escasd#@$Q@#@ESfsd@#fdsfwerq#dfsdf.sdfserew"

// ErrForbidden is returned when a auth issue is identified.
var ErrForbidden = errors.New("attempted action is not allowed")

// Claims represents the authorization claims transmitted via a JWT.
type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

type Auth struct {
	log    *logger.Logger
	method jwt.SigningMethod
	parser *jwt.Parser
	Issuer string
}

// Config represents information required to initialize auth.
type Config struct {
	Log    *logger.Logger
	Issuer string
}

func New(cfg Config) (*Auth, error) {
	return &Auth{
		log:    cfg.Log,
		method: jwt.GetSigningMethod(jwt.SigningMethodHS256.Name),
		parser: jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name})),
		Issuer: cfg.Issuer,
	}, nil
}

// GenerateToken generates a signed JWT token string representing the user Claims.
func (a *Auth) GenerateToken(kid string, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	// token.Header["kid"] = kid
	str, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return str, nil
}

// Authenticate processes the token to validate the sender's token is valid.
func (a *Auth) Authenticate(ctx context.Context, bearerToken string) (Claims, error) {
	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return Claims{}, errors.New("expected authorization header format: Bearer <token>")
	}

	var claims Claims

	token, err := a.parser.ParseWithClaims(parts[1], &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return Claims{}, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return claims, fmt.Errorf("token is not valid: %w", ErrForbidden)
	}

	if claims.Issuer != a.Issuer {
		return claims, fmt.Errorf("invalid token issuer: %w", ErrForbidden)
	}

	// Perform an extra level of authentication verification with OPA.

	// kidRaw, exists := token.Header["kid"]
	// if !exists {
	// 	return Claims{}, fmt.Errorf("kid missing from header: %w", err)
	// }

	// _, ok := kidRaw.(string)
	// if !ok {
	// 	return Claims{}, fmt.Errorf("kid malformed: %w", err)
	// }

	// Check the database for this user to verify they are still enabled.

	return claims, nil
}
