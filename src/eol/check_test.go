package eol

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"oduludo.io/eol/cfg"
	"os"
	"strings"
	"testing"
)

const (
	sourceBase               = "http://static/"
	unencryptedCustomSource1 = sourceBase + "example_datasource_readonly.json"
	unencryptedCustomSource2 = sourceBase + "example_datasource_readonly_2.json"
	encryptedCustomSource1   = sourceBase + "example_datasource_readonly_encrypted.json"
	encryptedCustomSource2   = sourceBase + "example_datasource_readonly_encrypted_2.json"
	encryptionKey1           = "rijltlTWRbYHtVaS"
	encryptionKey2           = "McjXXQNNLAircfSi"
)

const listVersionsTableHeader = `
+------------+--------------+
| Cycle name | Release name |
+------------+--------------+`

func performCheckCmdUnderTesting(t *testing.T, args ...string) (*bytes.Buffer, *cobra.Command, error) {
	if err := os.Setenv(cfg.IsIntegrationTestEnvKey, "true"); err != nil {
		t.Fatal(err)
	}

	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs(args)
	return actual, rootCmd, rootCmd.Execute()
}

// TestIntegrationCheckCmd performs an integration test, using the real datasource.
//
// Be mindful we don't control the data https://endoflife.date/ gives, so this test
// may break without any changes to the EOL codebase. If this integration test breaks
// it is likely the whole application can no longer fetch its data for usage.
func TestIntegrationCheckCmdWithoutStatusCode(t *testing.T) {
	actual, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"2.7",
	)

	if err != nil {
		t.Fatal(err)
	}

	expected := "true"

	assert.Equal(t, actual.String(), expected, "actual is not expected")

	if err := os.Unsetenv(cfg.IsIntegrationTestEnvKey); err != nil {
		t.Fatal(err)
	}
}

// Assert a non-existent version results in the table being printed.
// The exact contents don't matter.
func TestIntegrationCheckCmdVersionNotFound(t *testing.T) {
	actual, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"non-existent",
	)

	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, strings.Contains(actual.String(), listVersionsTableHeader))

	if err := os.Unsetenv(cfg.IsIntegrationTestEnvKey); err != nil {
		t.Fatal(err)
	}
}

// ==========================
// Custom datasources
// ==========================

/*
Unhappy flows
*/
// Test the check command gives an error if both --source and --xsource flags are set.
func TestCheckCmdWithBothSourceAndXsourceFlags(t *testing.T) {
	_, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"2.7",
		"--source=http://static/foo",
		"--xsource=http://static/foo",
	)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), cfg.SourceXsourceXorMsg)
}

// Test the check command gives an error if the number of keys provided is less than the number of (x)source URLs configured.
// Do this both for --source and --xsource
func TestCheckCmdWithInsufficientKeys(t *testing.T) {
	for _, flag := range []string{"source", "xsource"} {
		_, _, err := performCheckCmdUnderTesting(
			t,
			"check",
			"ruby",
			"2.7",
			fmt.Sprintf("--%s=http://static/foo,http://static/bar", flag),
			"--key=one_key",
		)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), cfg.InvalidKeysNumMsg)
	}
}

// Test the check command gives an error if the number of keys provided is more than the number of (x)source URLs configured.
// Do this both for --source and --xsource
func TestCheckCmdWithTooManyKeys(t *testing.T) {
	for _, flag := range []string{"source", "xsource"} {
		_, _, err := performCheckCmdUnderTesting(
			t,
			"check",
			"ruby",
			"2.7",
			fmt.Sprintf("--%s=http://static/foo,http://static/bar", flag),
			"--key=one_key,two_key,three_key",
		)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), cfg.InvalidKeysNumMsg)
	}
}

/*
Happy flows and Integration testing (on docker compose localhost)
*/
func TestCheckCmdWithRootSourceAndOneCustomSource(t *testing.T) {
	// The targeted version (ruby 3.2) is present both in the root source and in the custom source
	buffer, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"3.2",
		fmt.Sprintf("--source=%s", unencryptedCustomSource1),
	)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(buffer.String(), "false"))
}

func TestCheckCmdWithRootSourceAndCustomSources(t *testing.T) {
	// The targeted version (ruby 3.3) is only present in the second custom source
	buffer, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"3.3",
		fmt.Sprintf("--source=%s,%s", unencryptedCustomSource1, unencryptedCustomSource2),
	)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(buffer.String(), "false"))
}

func TestCheckCmdWithRootSourceAndEncryptedCustomSources(t *testing.T) {
	// The targeted version (ruby 3.3) is present only in the second custom source
	buffer, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"3.3",
		fmt.Sprintf("--source=%s,%s", encryptedCustomSource1, encryptedCustomSource2),
		fmt.Sprintf("--key=%s,%s", encryptionKey1, encryptionKey2),
	)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(buffer.String(), "false"))
}

// Use both encrypted and decrypted custom sources
func TestCheckCmdWithRootSourceAndMixedCustomSources(t *testing.T) {
	// The targeted version (ruby 3.3) is present only in the second custom source
	buffer, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"3.3",
		fmt.Sprintf("--source=%s,%s", unencryptedCustomSource1, encryptedCustomSource2),
		fmt.Sprintf("--key=_,%s", encryptionKey2),
	)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(buffer.String(), "false"))
}

func TestCheckCmdWithCustomOneSourceOnly(t *testing.T) {
	buffer, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"3.2",
		fmt.Sprintf("--xsource=%s", unencryptedCustomSource1),
	)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(buffer.String(), "false"))
}

func TestCheckCmdWithCustomSourcesOnly(t *testing.T) {
	// The targeted version (ruby 3.3) is present only in the second custom source
	buffer, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"3.3",
		fmt.Sprintf("--xsource=%s,%s", unencryptedCustomSource1, unencryptedCustomSource2),
	)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(buffer.String(), "false"))
}

func TestCheckCmdWithEncryptedCustomSourcesOnly(t *testing.T) {
	// The targeted version (ruby 3.3) is present only in the second custom source
	buffer, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"3.3",
		fmt.Sprintf("--xsource=%s,%s", encryptedCustomSource1, encryptedCustomSource2),
		fmt.Sprintf("--key=%s,%s", encryptionKey1, encryptionKey2),
	)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(buffer.String(), "false"))
}

// Use both encrypted and decrypted custom sources
func TestCheckCmdWithMixedCustomSourcesOnly(t *testing.T) {
	// The targeted version (ruby 3.3) is present only in the second custom source
	buffer, _, err := performCheckCmdUnderTesting(
		t,
		"check",
		"ruby",
		"3.3",
		fmt.Sprintf("--xsource=%s,%s", encryptedCustomSource1, unencryptedCustomSource2),
		fmt.Sprintf("--key=%s,_", encryptionKey1),
	)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(buffer.String(), "false"))
}
