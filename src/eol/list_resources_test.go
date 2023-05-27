package eol

import (
	"bytes"
	"github.com/stretchr/testify/assert"
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

func TestListResources(t *testing.T) {
	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"list-resources"})

	err := rootCmd.Execute()

	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, strings.Contains(actual.String(), resourcesText))
}

const filteredResourcesText = `Found 3 resources:
- go
- google-kubernetes-engine
- gorilla`

func TestFilteredListResources(t *testing.T) {
	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd(actual)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"list-resources", "--contains=go"})

	err := rootCmd.Execute()

	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, strings.Contains(actual.String(), filteredResourcesText))
}
