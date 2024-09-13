package message

import (
	"encoding/binary"

	"github.com/pkg/errors"
)

var _ IKEPayload = &KeyExchange{}

type KeyExchange struct {
	DiffieHellmanGroup uint16
	KeyExchangeData    []byte
}

func (keyExchange *KeyExchange) Type() IKEPayloadType { return TypeKE }

func (keyExchange *KeyExchange) marshal() ([]byte, error) {
	keyExchangeData := make([]byte, 4)
	binary.BigEndian.PutUint16(keyExchangeData[0:2], keyExchange.DiffieHellmanGroup)
	keyExchangeData = append(keyExchangeData, keyExchange.KeyExchangeData...)
	return keyExchangeData, nil
}

func (keyExchange *KeyExchange) unmarshal(b []byte) error {
	if len(b) > 0 {
		// bounds checking
		if len(b) <= 4 {
			return errors.Errorf("KeyExchange: No sufficient bytes to decode next key exchange data")
		}

		keyExchange.DiffieHellmanGroup = binary.BigEndian.Uint16(b[0:2])
		keyExchange.KeyExchangeData = append(keyExchange.KeyExchangeData, b[4:]...)
	}

	return nil
}
