package hyconutil

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test Blake2bHash with string params
func TestBlake2bHashString(t *testing.T) {
	input := "test black2b hash function"
	output := "b0f260dae1f5a49fdf7e5236b972d63d1e71525cac478bfb92261dcce722efb2"
	res, err := Blake2bHash(input)
	assert.Nil(t, err)
	resStr := hex.EncodeToString(res)
	assert.Equal(t, output, resStr, "The two should be same")
}

// test Blake2bHash with Byte Array params
func TestBlake2bHashByteArray(t *testing.T) {
	input := "test black2b hash function"
	output := "b0f260dae1f5a49fdf7e5236b972d63d1e71525cac478bfb92261dcce722efb2"
	inputByteArray := []uint8(input)
	res, err := Blake2bHash(inputByteArray)
	assert.Nil(t, err)
	resStr := hex.EncodeToString(res)
	assert.Equal(t, output, resStr, "The two should be same")
}

// test Blake2bHash with others types input
func TestBlake2bHashOthers(t *testing.T) {
	input := 234532
	_, err := Blake2bHash(input)
	assert.EqualError(t, err, "object type only can be string or []uint8")
}
