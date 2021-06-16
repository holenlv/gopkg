package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSet(t *testing.T) {
	{
		s := NewStringSet()
		s.Add("1")
		assert.Equal(t, s.Contains("1"), true)
		assert.Equal(t, s.Contains(""), false)
		assert.Equal(t, s.IsEmpty(), false)
		assert.Equal(t, s.Keys(), []string{"1"})

		s.Remove("1")
		assert.Equal(t, s.IsEmpty(), true)
		assert.Equal(t, s.Contains(""), false)
		assert.Equal(t, s.Keys(), []string{})
	}

	{
		s1 := NewStringSet().Add("1", "2", "3", "4")
		s1Copied := s1.Copy()
		s2 := NewStringSet().Add("3", "4", "5", "6")
		s2Copied := s2.Copy()
		union := s1.Union(s2)
		assert.Equal(t, union.IsEqual(NewStringSet().Add("1", "2", "3", "4", "5", "6")), true)
		intersection := s1.Intersection(s2)
		assert.Equal(t, intersection.IsEqual(NewStringSet().Add("3", "4")), true)
		assert.Equal(t, s1.IsEqual(s1Copied), true)
		assert.Equal(t, s2.IsEqual(s2Copied), true)
	}
}
