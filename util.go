package hyconutil

import (
	"errors"

	"golang.org/x/crypto/blake2b"
)

// make a blake2b hash on the string or buffer([]byte or []uint8)
// byte is alias of uint8
func Blake2bHash(ob interface{}) ([]uint8, error) {
	switch input := ob.(type) {
	case []uint8:
		hash := blake2b.Sum256(input)
		return hash[:], nil
	case string:
		hash := blake2b.Sum256([]uint8(input))
		return hash[:], nil
	default:
		return nil, errors.New("object type only can be string or []uint8")
	}
}
