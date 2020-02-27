package hyconutil

import (
	"errors"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/blake2b"
)

// make a blake2b hash on the string or buffer([]byte or []uint8)
// byte is alias of uint8
func Blake2bHash(ob []uint8) []uint8 {
	hash := blake2b.Sum256(ob)
	return hash[:]
}

func PublicKeyToAddress(publicKey []uint8) []uint8 {
	hash := Blake2bHash(publicKey)
	address := hash[12:]
	return address
}

func AddressCheckSum(arr []uint8) string {
	hash := Blake2bHash(arr)
	str := base58.Encode(hash)
	return str[:4]
}

func AddressToString(address []uint8) string {
	return "H" + base58.Encode(address) + AddressCheckSum(address)
}

func AddressToUint8Array(address string) ([]uint8, error) {
	if !strings.HasPrefix(address, "H") {
		return nil, errors.New("Address is invalid. Expected address to start with 'H'")
	}
	length := len(address)
	if length < 5 {
		return nil, errors.New("Address must be 20 bytes long")
	}
	check := address[length-4:]
	addr := address[1 : length-4]
	out := base58.Decode(addr)
	if len(out) != 20 {
		return nil, errors.New("Address must be 20 bytes long")
	}
	expectedChecksum := AddressCheckSum(out)
	if expectedChecksum != check {
		return nil, errors.New(`Address hash invalid checksum '` + check + `' expected '` + expectedChecksum + `'`)
	}
	return out, nil
}
