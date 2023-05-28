package eol

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func performCmdUnderTesting(args ...string) (*bytes.Buffer, *cobra.Command, error) {
	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs(args)
	return actual, rootCmd, rootCmd.Execute()
}

func testFileBeforeEncryption(t *testing.T, file string) {
	data, err := os.ReadFile(file)
	assert.Nil(t, err)

	bareText := strings.TrimSpace(string(data))
	firstChar := fmt.Sprintf("%c", bareText[0])
	lastChar := fmt.Sprintf("%c", bareText[len(bareText)-1])

	assert.Equal(t, "{", firstChar, "unexpected first token in file")
	assert.Equal(t, "}", lastChar, "unexpected last token in file")
}

func testFileAfterEncryption(t *testing.T, file string) {
	encryptedData, err := os.ReadFile(file)
	assert.Nil(t, err)

	encryptedText := strings.TrimSpace(string(encryptedData))
	firstChar := fmt.Sprintf("%c", encryptedText[0])
	lastChar := fmt.Sprintf("%c", encryptedText[len(encryptedText)-1])

	// The encrypted data should not contain any new lines and should not begin or end with '{' and '}' respectively
	assert.False(t, strings.Contains(string(encryptedData), "\n"))
	assert.NotEqual(t, "{", firstChar, "unexpected first token in encrypted data")
	assert.NotEqual(t, "}", lastChar, "unexpected last token in encrypted data")
}

// Perform `eol source encrypt FILE`
func TestEncryptCmdWithFile(t *testing.T) {
	inputFile := "/home/files/example_datasource_1.json"

	// Assert the situation pre-run
	testFileBeforeEncryption(t, inputFile)

	// Run the command
	buffer, _, err := performCmdUnderTesting(
		"source",
		"encrypt",
		inputFile,
	)
	assert.Nil(t, err)

	// Check the stdout output
	assert.True(t, strings.Contains(buffer.String(), "Generated new encryption key:"))
	assert.True(t, strings.Contains(buffer.String(), "Store it well!"))

	// Assert the situation post-run
	testFileAfterEncryption(t, inputFile)
}

// Perform `eol source encrypt FILE --to=OUTPUT`
func TestEncryptCmdWithFileAndOutput(t *testing.T) {
	inputFile := "/home/files/example_datasource_2.json"
	targetFile := "/home/files/encrypted_datasource_2.json"

	// Assert the situation pre-run
	testFileBeforeEncryption(t, inputFile)

	// Ensure the targetFile does not yet exist
	_, err := os.Stat(targetFile)
	assert.True(t, errors.Is(err, os.ErrNotExist))

	// Run the command
	buffer, _, err := performCmdUnderTesting(
		"source",
		"encrypt",
		inputFile,
		fmt.Sprintf("--to=%s", targetFile),
	)
	assert.Nil(t, err)

	// Check the stdout output
	assert.True(t, strings.Contains(buffer.String(), "Generated new encryption key:"))
	assert.True(t, strings.Contains(buffer.String(), "Store it well!"))

	// Assert the situation post-run
	// This will implicitly check the targetFile now exists, as the function will fail if it is missing
	testFileAfterEncryption(t, targetFile)
}

// Perform `eol source encrypt FILE --key=KEY`
func TestEncryptCmdWithFileAndKey(t *testing.T) {
	inputFile := "/home/files/example_datasource_3.json"
	key := "1234567890abcdef" // Must be 16 chars long

	// Assert the situation pre-run
	testFileBeforeEncryption(t, inputFile)

	// Run the command
	buffer, _, err := performCmdUnderTesting(
		"source",
		"encrypt",
		inputFile,
		fmt.Sprintf("--key=%s", key),
	)
	assert.Nil(t, err)

	// Check the stdout output
	assert.False(t, strings.Contains(buffer.String(), "Generated new encryption key:"))
	assert.False(t, strings.Contains(buffer.String(), key))
	assert.False(t, strings.Contains(buffer.String(), "Store it well!"))

	// Assert the situation post-run
	testFileAfterEncryption(t, inputFile)
}

// Perform `eol source encrypt FILE --to=OUTPUT --key=KEY`
func TestEncryptCmdWithFileAndOutputAndKey(t *testing.T) {
	inputFile := "/home/files/example_datasource_4.json"
	targetFile := "/home/files/encrypted_datasource_4.json"
	key := "1234567890abcdef" // Must be 16 chars long

	// Assert the situation pre-run
	testFileBeforeEncryption(t, inputFile)

	// Run the command
	buffer, _, err := performCmdUnderTesting(
		"source",
		"encrypt",
		inputFile,
		fmt.Sprintf("--key=%s", key),
		fmt.Sprintf("--to=%s", targetFile),
	)
	assert.Nil(t, err)

	// Check the stdout output
	assert.False(t, strings.Contains(buffer.String(), "Generated new encryption key:"))
	assert.False(t, strings.Contains(buffer.String(), key))
	assert.False(t, strings.Contains(buffer.String(), "Store it well!"))

	// Assert the situation post-run
	// This will implicitly check the targetFile now exists, as the function will fail if it is missing
	testFileAfterEncryption(t, targetFile)
}
