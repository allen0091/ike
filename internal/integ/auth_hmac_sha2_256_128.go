package integ

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"

	"github.com/free5gc/ike/types"
)

const string_AUTH_HMAC_SHA2_256_128 string = "AUTH_HMAC_SHA2_256_128"

func toString_AUTH_HMAC_SHA2_256_128(attrType uint16, intValue uint16, bytesValue []byte) string {
	return string_AUTH_HMAC_SHA2_256_128
}

var (
	_ INTEGType  = &AUTH_HMAC_SHA2_256_128{}
	_ INTEGKType = &AUTH_HMAC_SHA2_256_128{}
)

type AUTH_HMAC_SHA2_256_128 struct {
	priority     uint32
	keyLength    int
	outputLength int
}

func (t *AUTH_HMAC_SHA2_256_128) TransformID() uint16 {
	return types.AUTH_HMAC_SHA2_256_128
}

func (t *AUTH_HMAC_SHA2_256_128) getAttribute() (bool, uint16, uint16, []byte) {
	return false, 0, 0, nil
}

func (t *AUTH_HMAC_SHA2_256_128) setPriority(priority uint32) {
	t.priority = priority
}

func (t *AUTH_HMAC_SHA2_256_128) Priority() uint32 {
	return t.priority
}

func (t *AUTH_HMAC_SHA2_256_128) GetKeyLength() int {
	return t.keyLength
}

func (t *AUTH_HMAC_SHA2_256_128) GetOutputLength() int {
	return t.outputLength
}

func (t *AUTH_HMAC_SHA2_256_128) Init(key []byte) hash.Hash {
	if len(key) == 32 {
		return hmac.New(sha256.New, key)
	} else {
		return nil
	}
}

func (t *AUTH_HMAC_SHA2_256_128) XFRMString() string {
	return "hmac(sha256)"
}
