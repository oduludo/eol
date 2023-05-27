package eol

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"oduludo.io/eol/cfg"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const listVersionsTableHeader = `
+------------+--------------+
| Cycle name | Release name |
+------------+--------------+`

// TestIntegrationCheckCmd performs an integration test, using the real datasource.
//
// Be mindful we don't control the data https://endoflife.date/ gives, so this test
// may break without any changes to the EOL codebase. If this integration test breaks
// it is likely the whole application can no longer fetch its data for usage.
func TestIntegrationCheckCmdWithoutStatusCode(t *testing.T) {
	if err := os.Setenv(cfg.IsIntegrationTestEnvKey, "true"); err != nil {
		t.Fatal(err)
	}

	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"check", "ruby", "2.7"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}

	expected := "true"

	assert.Equal(t, actual.String(), expected, "actual is not expected")

	if err := os.Unsetenv(cfg.IsIntegrationTestEnvKey); err != nil {
		t.Fatal(err)
	}
}

func TestIntegrationCheckCmdWithStatusCode(t *testing.T) {
	if err := os.Setenv(cfg.IsIntegrationTestEnvKey, "true"); err != nil {
		t.Fatal(err)
	}

	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"check", "ruby", "2.7"})

	err := rootCmd.Execute()
	if e, ok := err.(*exec.ExitError); ok && e.ExitCode() != eolExitCode {
	}

	expected := "true"

	assert.Equal(t, actual.String(), expected, "actual is not expected")

	if err := os.Unsetenv(cfg.IsIntegrationTestEnvKey); err != nil {
		t.Fatal(err)
	}
}

// Assert a non-existent version results in the table being printed.
// The exact contents don't matter.
func TestIntefrationCheckCmdVersionNotFound(t *testing.T) {
	if err := os.Setenv(cfg.IsIntegrationTestEnvKey, "true"); err != nil {
		t.Fatal(err)
	}

	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"check", "ruby", "non-existent"})

	err := rootCmd.Execute()

	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, strings.Contains(actual.String(), listVersionsTableHeader))

	if err := os.Unsetenv(cfg.IsIntegrationTestEnvKey); err != nil {
		t.Fatal(err)
	}
}
