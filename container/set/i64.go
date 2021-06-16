package set

type I64Set struct {
	data map[int64]struct{}
}

func NewI64Set() *I64Set {
	return &I64Set{make(map[int64]struct{})}
}

func (s *I64Set) Add(values ...int64) *I64Set {
	for _, v := range values {
		if !s.Contains(v) {
			s.data[v] = struct{}{}
		}
	}
	return s
}

func (s *I64Set) Remove(values ...int64) *I64Set {
	for _, v := range values {
		if s.Contains(v) {
			delete(s.data, v)
		}
	}
	return s
}

func (s *I64Set) Contains(val int64) bool {
	_, ok := s.data[val]
	return ok
}

func (s *I64Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *I64Set) Keys() []int64 {
	keys := make([]int64, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *I64Set) Size() int {
	if s == nil {
		return 0
	}
	return len(s.data)
}

func (s *I64Set) Copy() *I64Set {
	copied := NewI64Set()
	if s == nil || s.IsEmpty() {
		return copied
	}
	for k := range s.data {
		copied.data[k] = struct{}{}
	}
	return copied
}

func (s *I64Set) IsEqual(as *I64Set) bool {
	if s.Size() != as.Size() {
		return false
	}
	for k := range s.data {
		if !as.Contains(k) {
			return false
		}
	}
	return true
}

func (s *I64Set) Intersection(as *I64Set) *I64Set {
	ret := NewI64Set()
	for k := range s.data {
		if as.Contains(k) {
			ret.Add(k)
		}
	}
	return ret
}

func (s *I64Set) Union(as *I64Set) *I64Set {
	ret := s.Copy()
	if as == nil || as.IsEmpty() {
		return ret
	}
	keys := as.Keys()
	return ret.Add(keys...)
}
