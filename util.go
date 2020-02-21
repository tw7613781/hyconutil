package hyconutil

import (
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
