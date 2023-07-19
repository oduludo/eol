package eol

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"oduludo.io/eol/cfg"
	"oduludo.io/eol/eol/utils"
	"oduludo.io/eol/pkg/crypt"
	"strings"
	"testing"
)

// Perform `eol source verify FILE`
func TestVerifyCmdWithFile(t *testing.T) {
	inputFile := "/home/files/example_datasource_readonly.json"

	// Run the command
	buffer, _, err := performCmdUnderTesting(
		"source",
		"verify",
		inputFile,
	)
	assert.Nil(t, err)

	// Check the stdout output
	assert.True(t, strings.Contains(buffer.String(), cfg.DatasourceValidMsg))
}

// Perform `eol source verify FILE` on an invalid file
func TestVerifyCmdWithInvalidFile(t *testing.T) {
	inputFile := "/home/files/invalid_datasource_1.json"

	// Run the command
	_, _, err := performCmdUnderTesting(
		"source",
		"verify",
		inputFile,
	)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), cfg.DatasourceInvalidMsg)
}

// Perform `eol source verify FILE --key=KEY`
func TestVerifyCmdWithEncryptedFile(t *testing.T) {
	encryptionInputFile := "/home/files/example_datasource_10.json"
	encryptionOutputFile := "/home/files/encrypted_datasource_10.json"

	// Encrypt the input file
	newKey := crypt.GenerateKey()
	key, _ := prepareForDecryptionWithOutputAndKey(t, encryptionInputFile, encryptionOutputFile, newKey)
	assert.Equal(t, newKey, key)

	err, isEncrypted := utils.FileIsEncrypted(encryptionOutputFile)
	assert.Nil(t, err)
	assert.True(t, isEncrypted)

	// Run the command
	buffer, _, err := performCmdUnderTesting(
		"source",
		"verify",
		encryptionOutputFile,
		fmt.Sprintf("--key=%s", key),
	)
	assert.Nil(t, err)

	// Check the stdout output
	assert.True(t, strings.Contains(buffer.String(), cfg.DatasourceValidMsg))
}

// Perform `eol source verify FILE --key=KEY` on invalid file
func TestVerifyCmdWithEncryptedInvalidFile(t *testing.T) {
	encryptionInputFile := "/home/files/invalid_datasource_2.json"
	encryptionOutputFile := "/home/files/encrypted_invalid_datasource_2.json"

	// Encrypt the input file
	newKey := crypt.GenerateKey()
	key, _ := prepareForDecryptionWithOutputAndKey(t, encryptionInputFile, encryptionOutputFile, newKey)
	assert.Equal(t, newKey, key)

	err, isEncrypted := utils.FileIsEncrypted(encryptionOutputFile)
	assert.Nil(t, err)
	assert.True(t, isEncrypted)

	// Run the command
	_, _, err = performCmdUnderTesting(
		"source",
		"verify",
		encryptionOutputFile,
		fmt.Sprintf("--key=%s", key),
	)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), cfg.DatasourceInvalidMsg)
}
