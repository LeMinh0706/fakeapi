package token

import (
	"time"
)

type Maker interface {
	CreateToken(sub int32, duration time.Duration) (string, error)
	VerifyToken(tokenStr string) (*Payload, error)
}
