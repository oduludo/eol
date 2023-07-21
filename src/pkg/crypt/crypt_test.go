package crypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptionDecryptionCycle(t *testing.T) {
	text := "Foobar_125!"
	key := GenerateKey()

	// Encrypt and ensure it's different from the source text
	encrypted, err := Encrypt(text, key)

	assert.Nil(t, err)
	assert.NotEqual(t, "", encrypted)
	assert.NotEqual(t, text, encrypted)

	// Decrypt and ensure it matches the source text
	decrypted, err := Decrypt(encrypted, key)
	assert.Nil(t, err)
	assert.Equal(t, text, decrypted)
}
