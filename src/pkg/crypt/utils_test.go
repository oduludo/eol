package crypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	// Check the length is consistently correct
	for i := 0; i < 10; i++ {
		assert.Equal(t, EncryptionKeyLength, len(GenerateKey()))
	}

	// Check the function does not return the same result
	assert.NotEqual(t, GenerateKey(), GenerateKey())
}

func TestValidateValidKey(t *testing.T) {
	assert.True(t, ValidateKey(GenerateKey()))
}

func TestValidateInvalidKey(t *testing.T) {
	assert.False(t, ValidateKey("nonsense"))
}
