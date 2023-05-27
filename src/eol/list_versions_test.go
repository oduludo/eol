package eol

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const listVersionsTable = `
+------------+--------------+
| Cycle name | Release name |
+------------+--------------+
|    3.2     |              |
|    3.1     |              |
|    3.0     |              |
|    2.7     |              |
|    2.6     |              |
|    2.5     |              |
|    2.4     |              |
|    2.3     |              |
|    2.2     |              |
|    2.1     |              |
|   2.0.0    |              |
|   1.9.3    |              |
+------------+--------------+`

// TestListVersions checks the exact table is written to the buffer.
func TestListVersions(t *testing.T) {
	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"list-versions", "ruby"})

	err := rootCmd.Execute()

	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, strings.Contains(actual.String(), listVersionsTable))
}
