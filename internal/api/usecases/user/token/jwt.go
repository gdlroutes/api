package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gdlroutes/api/internal/api/models"
	"github.com/gdlroutes/api/internal/api/usecases/user"
)

type jwtTokenGenerator struct {
	tokenTTL time.Duration
	key      []byte
}

var _ user.TokenGenerator = &jwtTokenGenerator{}

// NewJWT creates a new JWT token generator
func NewJWT(tokenTTL time.Duration, key string) (user.TokenGenerator, error) {
	return &jwtTokenGenerator{
		tokenTTL: tokenTTL,
		key:      []byte(key),
	}, nil
}

func (g *jwtTokenGenerator) GenerateToken(userID string, start time.Time) (*models.Token, error) {
	expiresOn := start.Add(g.tokenTTL)
	claims := jwt.MapClaims(map[string]interface{}{
		"sub": userID,
		"exp": expiresOn.Unix(),
	})

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString(g.key)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Name:    "Access",
		Token:   signedToken,
		Expires: expiresOn,
	}, nil
}
