package token

import (
	"errors"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var ErrInvalid = errors.New("invalid token")

// CreateToken implements Maker.
func (p *PasetoMaker) CreateToken(sub int32, duration time.Duration) (string, error) {
	payload, err := NewPayload(uuid.New(), sub, duration)
	if err != nil {
		return "", err
	}
	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}

// VerifyToken implements Maker.
func (p *PasetoMaker) VerifyToken(tokenStr string) (*Payload, error) {
	payload := &Payload{}
	err := p.paseto.Decrypt(tokenStr, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalid
	}
	return payload, nil
}

func NewPasetoMaker(symmetricKey []byte) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid symmetric key size, must be 32 bytes")
	}
	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: symmetricKey,
	}, nil
}
