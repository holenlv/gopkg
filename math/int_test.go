package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestI(t *testing.T) {
	assert.Equal(t, MaxI(2, 1), 2)
	assert.Equal(t, MaxI(0, 1), 1)
	assert.Equal(t, MinI(-1, 1), -1)
	assert.Equal(t, MinI(0, 1), 0)
	assert.Equal(t, AbsI(1), 1)
	assert.Equal(t, AbsI(-1), 1)
	assert.Equal(t, AbsI(0), 0)
}
