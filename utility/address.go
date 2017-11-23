package utility

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
)

// ReverseByteSlice reverts content of given slice in-place
func ReverseByteSlice(b []byte) {
	for i := len(b)/2 - 1; i >= 0; i-- {
		opp := len(b) - 1 - i
		b[i], b[opp] = b[opp], b[i]
	}
}

// AddressToHash converts base58 Neo address into little endian hash160 format
func AddressToHash(addr string, addrVer byte) (string, error) {
	b := NewBase58()
	data, err := b.Decode(addr)
	if err != nil {
		return "", err
	}
	if len(data) != 25 {
		return "", errors.New("Wrong address length")
	}
	if data[0] != addrVer {
		return "", errors.New("Incorrect coin version")
	}

	checksum1stRound := sha256.Sum256(data[:21])
	checksum2ndRound := sha256.Sum256(checksum1stRound[:])
	checksumTail := checksum2ndRound[:4]
	if !bytes.Equal(data[21:], checksumTail) {
		return "", errors.New("Address format error")
	}
	hashBytes := data[1:21]
	ReverseByteSlice(hashBytes)
	// do not attempt to use `data` variable beyond this point
	// operations on slice directly mutate underlying content

	return fmt.Sprintf("%x", hashBytes), nil
}
