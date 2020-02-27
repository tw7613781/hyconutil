package hyconutil

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = "test black2b hash function"
var inputBuf = []uint8(input)
var hash []uint8
var address []uint8
var addrHycon string

// test Blake2bHash with Byte Array params
func TestBlake2bHash(t *testing.T) {
	output := "b0f260dae1f5a49fdf7e5236b972d63d1e71525cac478bfb92261dcce722efb2"
	hash = Blake2bHash(inputBuf)
	resStr := hex.EncodeToString(hash)
	assert.Equal(t, 32, len(hash), "The hash length is 32 bytes")
	assert.Equal(t, output, resStr, "The two should be same")
}

func TestPublicKeyToAddress(t *testing.T) {
	output := "3ef61acd9b8649563835252cb7068e34312ca7a4"
	address = PublicKeyToAddress(hash)
	resStr := hex.EncodeToString(address)
	assert.Equal(t, output, resStr, "The two should be same")
}

func TestAddressCheckSum(t *testing.T) {
	output := "2LmU"
	res := AddressCheckSum(address)
	assert.Equal(t, output, res, "The two should be same")
}

func TestAddressToString(t *testing.T) {
	addrHycon = "HsskvzK6pzYiYnMoJkNhPGjc8fP92LmU"
	res := AddressToString(address)
	assert.Equal(t, addrHycon, res, "The two should be same")
}

func TestAddressToUint8Array(t *testing.T) {
	output := []uint8{62,
		246,
		26,
		205,
		155,
		134,
		73,
		86,
		56,
		53,
		37,
		44,
		183,
		6,
		142,
		52,
		49,
		44,
		167,
		164}
	res, err := AddressToUint8Array(addrHycon)
	assert.NoError(t, err)
	assert.Equal(t, output, res, "The two should be same")
}

func TestAddressToStringWithIncorrectCheckSum(t *testing.T) {
	output := "HsskvzK6pzYiYnMoJkNhPGjc8fP92LmA"
	_, err := AddressToUint8Array(output)
	assert.EqualError(t, err, `Address hash invalid checksum '2LmA' expected '2LmU'`)
}

func TestAddressToStringWithIncorrectStartingAlphabet(t *testing.T) {
	output := "AsskvzK6pzYiYnMoJkNhPGjc8fP92LmU"
	_, err := AddressToUint8Array(output)
	assert.EqualError(t, err, `Address is invalid. Expected address to start with 'H'`)
}

func TestAddressToStringWithIncorrectLength(t *testing.T) {
	output := "Hss"
	_, err := AddressToUint8Array(output)
	assert.EqualError(t, err, "Address must be 20 bytes long")
}
