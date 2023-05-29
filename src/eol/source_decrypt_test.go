package eol

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"oduludo.io/eol/eol/utils"
	"oduludo.io/eol/pkg/crypt"
	"testing"
)

func prepareForDecryptionWithOutputAndKey(t *testing.T, file string, output string, key string) (string, error) {
	// Start by encrypting a file
	args := []string{
		"source",
		"encrypt",
		file,
	}

	if output != "" {
		args = append(args, fmt.Sprintf("--to=%s", output))
	}

	if key == "" {
		args = append(args, "--x-only-print-key")
	} else {
		args = append(args, fmt.Sprintf("--key=%s", key))
	}

	actual, _, err := performCmdUnderTesting(args...)
	assert.Nil(t, err)

	if key == "" {
		// No key provided
		// Extract the generated key from buffer, as we need it for decryption
		// Due to usage of hidden flag --x-only-print-key, all that's in buffer is the key
		key = actual.String()
	}
	assert.True(t, crypt.ValidateKey(key))

	return key, err
}

func prepareForDecryption(t *testing.T, file string) (string, error) {
	return prepareForDecryptionWithOutputAndKey(t, file, "", "")
}

// Run `eol source decrypt FILE`
func TestDecryptionCmdWithFile(t *testing.T) {
	inputFile := "/home/files/example_datasource_5.json"
	key, _ := prepareForDecryption(t, inputFile)

	// Now it's time to test the decryption
	_, _, err := performCmdUnderTesting(
		"source",
		"decrypt",
		inputFile,
		fmt.Sprintf("--key=%s", key),
	)
	assert.Nil(t, err)

	// Check the same file now contains JSON data
	err, ok := utils.FileContainsJson(inputFile)
	assert.Nil(t, err)
	assert.True(t, ok, "file does not contain JSON data")
}

// Run `eol source decrypt FILE --to=OUTPUT`
func TestDecryptionCmdWithFileAndOutput(t *testing.T) {
	// The input file for ENCRYPTION will become the output file for DECRYPTION
	encryptionInputFile := "/home/files/example_datasource_6.json"
	// The output file for ENCRYPTION will become the input file for DECRYPTION
	encryptionOutputFile := "/home/files/encrypted_datasource_6.json"
	key, _ := prepareForDecryptionWithOutputAndKey(t, encryptionInputFile, encryptionOutputFile, "")

	// Now it's time to test the decryption
	_, _, err := performCmdUnderTesting(
		"source",
		"decrypt",
		encryptionOutputFile,
		fmt.Sprintf("--key=%s", key),
		fmt.Sprintf("--to=%s", encryptionInputFile),
	)
	assert.Nil(t, err)

	// Check the same file now contains JSON data
	err, ok := utils.FileContainsJson(encryptionInputFile)
	assert.Nil(t, err)
	assert.True(t, ok, "file does not contain JSON data")
}

// Run `eol source decrypt FILE --to=OUTPUT --key=KEY`
// Largely the same as TestDecryptionCmdWithFileAndOutput, but a key is provided for encrypting the data rather
// than the system coming up with a key in the fly.
func TestDecryptionCmdWithFileAndOutputAndKey(t *testing.T) {
	// The input file for ENCRYPTION will become the output file for DECRYPTION
	encryptionInputFile := "/home/files/example_datasource_7.json"
	// The output file for ENCRYPTION will become the input file for DECRYPTION
	encryptionOutputFile := "/home/files/encrypted_datasource_7.json"

	newKey := crypt.GenerateKey()
	key, _ := prepareForDecryptionWithOutputAndKey(t, encryptionInputFile, encryptionOutputFile, newKey)
	assert.Equal(t, newKey, key)

	// Now it's time to test the decryption
	_, _, err := performCmdUnderTesting(
		"source",
		"decrypt",
		encryptionOutputFile,
		fmt.Sprintf("--key=%s", key),
		fmt.Sprintf("--to=%s", encryptionInputFile),
	)
	assert.Nil(t, err)

	// Check the same file now contains JSON data
	err, ok := utils.FileContainsJson(encryptionInputFile)
	assert.Nil(t, err)
	assert.True(t, ok, "file does not contain JSON data")
}

// Run `eol source decrypt FILE --key=KEY`
func TestDecryptionCmdWithFileAndKey(t *testing.T) {
	file := "/home/files/example_datasource_8.json"

	newKey := crypt.GenerateKey()
	key, _ := prepareForDecryptionWithOutputAndKey(t, file, "", newKey)
	assert.Equal(t, newKey, key)

	// Now it's time to test the decryption
	_, _, err := performCmdUnderTesting(
		"source",
		"decrypt",
		file,
		fmt.Sprintf("--key=%s", key),
	)
	assert.Nil(t, err)

	// Check the same file now contains JSON data
	err, ok := utils.FileContainsJson(file)
	assert.Nil(t, err)
	assert.True(t, ok, "file does not contain JSON data")
}
