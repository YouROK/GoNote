package utils

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
)

func RandStr(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func Sanitize(s string) string {
	// Все символы кроме букв, цифр, пробелов и подчёркиваний → тире
	re := regexp.MustCompile(`[ _]+|[^a-zA-Z0-9-]+`)
	s = re.ReplaceAllString(s, "-")

	// Убираем повторяющиеся тире
	re2 := regexp.MustCompile("-+")
	s = re2.ReplaceAllString(s, "-")

	// Убираем тире в начале и конце
	s = strings.Trim(s, "-")

	return s
}
