package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	jwt.RegisteredClaims
}

func NewPayload(id uuid.UUID, sub uuid.UUID, duration time.Duration) (*Payload, error) {
	return &Payload{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id.String(),
			Subject:   sub.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}
