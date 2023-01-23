package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Base64Decode
func Base64Decode(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(hex.EncodeToString(data))
}

// TxHashFromTxData
func TxHashFromTxData(txData []byte) string {
	bd, err := Base64Decode(txData)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(bd)
	return hex.EncodeToString(hash[:])
}
