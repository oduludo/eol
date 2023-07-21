package crypt

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const EncryptionKeyLength = 16
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randStringBytesMask(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func GenerateKey() string {
	return randStringBytesMask(EncryptionKeyLength)
}

func ValidateKey(key string) bool {
	return len(key) == EncryptionKeyLength
}

func StringIsEncrypted(data string) (bool, error) {
	encryptedText := strings.TrimSpace(data)
	firstChar := fmt.Sprintf("%c", encryptedText[0])
	lastChar := fmt.Sprintf("%c", encryptedText[len(encryptedText)-1])

	// The encrypted data should not contain any new lines and should not begin or end with '{' and '}' respectively
	if strings.Contains(data, "\n") {
		return false, nil
	}

	if firstChar == "{" {
		return false, nil
	}

	if lastChar == "}" {
		return false, nil
	}

	return true, nil
}
