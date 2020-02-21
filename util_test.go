package hyconutil

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = "test black2b hash function"
var inputBuf = []uint8(input)

// test Blake2bHash with Byte Array params
func TestBlake2bHash(t *testing.T) {
	output := "b0f260dae1f5a49fdf7e5236b972d63d1e71525cac478bfb92261dcce722efb2"
	res := Blake2bHash(inputBuf)
	resStr := hex.EncodeToString(res)
	assert.Equal(t, output, resStr, "The two should be same")
}

func TestPublicKeyToAddress(t *testing.T) {
	output := "b972d63d1e71525cac478bfb92261dcce722efb2"
	res := PublicKeyToAddress(inputBuf)
	resStr := hex.EncodeToString(res)
	assert.Equal(t, output, resStr, "The two should be same")
}

func TestAddressCheckSum(t *testing.T) {
	output := "Cuj7"
	res := AddressCheckSum(inputBuf)
	assert.Equal(t, output, res, "The two should be same")
}

func TestAddressToString(t *testing.T) {
	output := "H4ZisVWsFYsBsXcgq8XYBLyvNFuarDEjqKuHoCuj7"
	res := AddressToString(inputBuf)
	assert.Equal(t, output, res, "The two should be same")
}
