package eol

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

// TestIntegrationCheckCmd performs an integration test, using the real datasource.
//
// Be mindful we don't control the data https://endoflife.date/ gives, so this test
// may break without any changes to the EOL codebase. If this integration test breaks
// it is likely the whole application can no longer fetch its data for usage.
func TestIntegrationCheckCmdWithoutStatusCode(t *testing.T) {
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
}

func TestIntegrationCheckCmdWithStatusCode(t *testing.T) {
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
}
