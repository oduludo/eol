package eol

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"oduludo.io/eol/cfg"
	"os"
	"strings"
	"testing"
)

const resourcesText = `Found 6 resources:
- go
- google-kubernetes-engine
- gorilla
- mysql
- neo4j
- ruby`

func performListResourcesCmdUnderTesting(t *testing.T, args ...string) (*bytes.Buffer, *cobra.Command, error) {
	if err := os.Setenv(cfg.IsIntegrationTestEnvKey, "false"); err != nil {
		t.Fatal(err)
	}

	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs(args)
	return actual, rootCmd, rootCmd.Execute()
}

func TestListResources(t *testing.T) {
	actual, _, _ := performListResourcesCmdUnderTesting(t, "list-resources")

	assert.True(t, strings.Contains(actual.String(), resourcesText))
}

const filteredResourcesText = `Found 3 resources:
- go
- google-kubernetes-engine
- gorilla`

func TestFilteredListResources(t *testing.T) {
	actual, _, _ := performListResourcesCmdUnderTesting(t, "list-resources", "--contains=go")

	assert.True(t, strings.Contains(actual.String(), filteredResourcesText))
}
