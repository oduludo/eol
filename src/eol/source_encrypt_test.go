package eol

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"oduludo.io/eol/eol/utils"
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
	err, ok := utils.FileContainsJson(file)
	assert.Nil(t, err)
	assert.True(t, ok, "file does not contain JSON data")
}

func testFileAfterEncryption(t *testing.T, file string) {
	err, ok := utils.FileIsEncrypted(file)
	assert.Nil(t, err)
	assert.True(t, ok, "file contains JSON data")
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

func TestXOnlyPrintKeyFlagIsHiddenInHelpMsg(t *testing.T) {
	buffer, _, err := performCmdUnderTesting(
		"source",
		"encrypt",
		"--help",
	)
	assert.Nil(t, err)

	assert.False(t, strings.Contains(buffer.String(), "x-only-print-key"))
}

func TestXOnlyPrintKeyOnlyPrintsKey(t *testing.T) {
	inputFile := "/home/files/example_datasource_9.json"
	buffer, _, err := performCmdUnderTesting(
		"source",
		"encrypt",
		inputFile,
		"--x-only-print-key",
	)
	assert.Nil(t, err)
	assert.Equal(t, 16, len(strings.TrimSpace(buffer.String())))
}
