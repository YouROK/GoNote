package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func RandStr(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return hex.EncodeToString(b)
}
