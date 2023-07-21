package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZip(t *testing.T) {
	arrayA := []int{2, 3, 4}
	arrayB := []string{"foo", "bar", "baz"}

	result, ok := Zip[int, string](arrayA, arrayB)

	assert.True(t, ok)

	assert.Equal(t, 2, result[0].A)
	assert.Equal(t, "foo", result[0].B)

	assert.Equal(t, 3, result[1].A)
	assert.Equal(t, "bar", result[1].B)

	assert.Equal(t, 4, result[2].A)
	assert.Equal(t, "baz", result[2].B)
}

func TestZipOnUnequallyLongArrays(t *testing.T) {
	arrayA := []int{2, 3}
	arrayB := []string{"foo", "bar", "baz"}

	result, ok := Zip[int, string](arrayA, arrayB)

	assert.False(t, ok)
	assert.Equal(t, 0, len(result))
}
