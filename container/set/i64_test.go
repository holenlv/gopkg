package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestI64Set(t *testing.T) {
	{
		s := NewI64Set()
		s.Add(1)
		assert.Equal(t, s.Contains(1), true)
		assert.Equal(t, s.Contains(0), false)
		assert.Equal(t, s.IsEmpty(), false)
		assert.Equal(t, s.Keys(), []int64{1})

		s.Remove(1)
		assert.Equal(t, s.IsEmpty(), true)
		assert.Equal(t, s.Contains(1), false)
		assert.Equal(t, s.Keys(), []int64{})
	}

	{
		s1 := NewI64Set().Add(1, 2, 3, 4)
		s1Copied := s1.Copy()
		s2 := NewI64Set().Add(3, 4, 5, 6)
		s2Copied := s2.Copy()
		union := s1.Union(s2)
		assert.Equal(t, union.IsEqual(NewI64Set().Add(1, 2, 3, 4, 5, 6)), true)
		intersection := s1.Intersection(s2)
		assert.Equal(t, intersection.IsEqual(NewI64Set().Add(3, 4)), true)
		assert.Equal(t, s1.IsEqual(s1Copied), true)
		assert.Equal(t, s2.IsEqual(s2Copied), true)
	}
}
