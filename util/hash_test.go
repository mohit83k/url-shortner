package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortKey_Consistency(t *testing.T) {
	url := "https://example.com"
	key1 := GenerateShortKey(url)
	key2 := GenerateShortKey(url)
	assert.Equal(t, key1, key2)
	assert.Len(t, key1, 6)
}
