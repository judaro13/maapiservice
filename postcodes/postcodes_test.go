package postcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitInChuncks(t *testing.T) {
	data := make([][]string, 30)
	result := splitInChuncks(data)
	assert.Equal(t, len(result), 1)

	data = make([][]string, 290)
	result = splitInChuncks(data)
	assert.Equal(t, len(result), 3)

}
