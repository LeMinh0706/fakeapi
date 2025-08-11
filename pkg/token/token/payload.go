package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	UserID int32 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewPayload(id uuid.UUID, sub int32, duration time.Duration) (*Payload, error) {
	return &Payload{
		UserID: sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}
