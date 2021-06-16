package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestI64(t *testing.T) {
	assert.Equal(t, MaxI64(2, 1), int64(2))
	assert.Equal(t, MaxI64(0, 1), int64(1))
	assert.Equal(t, MinI64(-1, 1), int64(-1))
	assert.Equal(t, MinI64(0, 1), int64(0))
	assert.Equal(t, AbsI64(1), int64(1))
	assert.Equal(t, AbsI64(-1), int64(1))
	assert.Equal(t, AbsI64(0), int64(0))
}
