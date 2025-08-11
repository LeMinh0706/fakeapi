package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(sub uuid.UUID, duration time.Duration) (string, error)
	VerifyToken(tokenStr string) (*Payload, error)
}
