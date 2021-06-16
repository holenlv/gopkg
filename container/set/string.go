package set

type StringSet struct {
	data map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{make(map[string]struct{})}
}

func (s *StringSet) Add(values ...string) *StringSet {
	for _, v := range values {
		if !s.Contains(v) {
			s.data[v] = struct{}{}
		}
	}
	return s
}

func (s *StringSet) Remove(values ...string) *StringSet {
	for _, v := range values {
		if s.Contains(v) {
			delete(s.data, v)
		}
	}
	return s
}

func (s *StringSet) Contains(val string) bool {
	_, ok := s.data[val]
	return ok
}

func (s *StringSet) IsEmpty() bool {
	return s.Size() == 0
}

func (s *StringSet) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *StringSet) Size() int {
	if s == nil {
		return 0
	}
	return len(s.data)
}

func (s *StringSet) Copy() *StringSet {
	copied := NewStringSet()
	if s == nil || s.IsEmpty() {
		return copied
	}
	for k := range s.data {
		copied.data[k] = struct{}{}
	}
	return copied
}

func (s *StringSet) IsEqual(as *StringSet) bool {
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

func (s *StringSet) Intersection(as *StringSet) *StringSet {
	ret := NewStringSet()
	for k := range s.data {
		if as.Contains(k) {
			ret.Add(k)
		}
	}
	return ret
}

func (s *StringSet) Union(as *StringSet) *StringSet {
	ret := s.Copy()
	if as == nil || as.IsEmpty() {
		return ret
	}
	keys := as.Keys()
	return ret.Add(keys...)
}
