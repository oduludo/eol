package eol

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestIntegrationCheckCmd performs an integration test, using the real datasource.
//
// Be mindful we don't control the data https://endoflife.date/ gives, so this test
// may break without any changes to the EOL codebase. If this integration test breaks
// it is likely the whole application can no longer fetch its data for usage.
func TestIntegrationCheckCmd(t *testing.T) {
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
