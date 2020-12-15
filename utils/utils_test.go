package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	uuid, err := NewUUID()
	assert.NoError(t, err)
	uuid2, err := NewUUID()
	assert.NoError(t, err)
	assert.NotEqual(t, uuid, uuid2)
}

func TestIsValidUUID(t *testing.T) {
	uuid, err := NewUUID()
	assert.NoError(t, err)
	assert.True(t, IsValidUUID(uuid))
	assert.False(t, IsValidUUID("uuid"))
}
