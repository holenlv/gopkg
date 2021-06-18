package convetor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, S2I64(""), int64(0))
	assert.Equal(t, S2I64("1"), int64(1))
	assert.Equal(t, S2I64("-1"), int64(-1))

	assert.Equal(t, I642S(0), "0")
	assert.Equal(t, I642S(1), "1")
	assert.Equal(t, I642S(-1), "-1")

	assert.Equal(t, I322S(0), "0")
	assert.Equal(t, I322S(1), "1")
	assert.Equal(t, I322S(-1), "-1")

	assert.Equal(t, Bytes2Str([]byte{'h', 'e', 'l', 'l', 'o'}), "hello")
	assert.Equal(t, Str2Bytes("hello"), []byte{'h', 'e', 'l', 'l', 'o'})
}
