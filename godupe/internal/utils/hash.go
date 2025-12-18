package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func CalculateHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashString := hex.EncodeToString(hash.Sum(nil))

	return hashString, nil
}
