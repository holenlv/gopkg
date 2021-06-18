package convetor

import (
	"strconv"
	"unsafe"
)

func S2I64(s string) int64 {
	i64, _ := strconv.ParseInt(s, 10, 64)
	return i64
}

func S2UI32(s string) uint32 {
	ui32, _ := strconv.ParseUint(s, 10, 32)
	return uint32(ui32)
}

func S2I32(s string) int32 {
	i32, _ := strconv.ParseInt(s, 10, 32)
	return int32(i32)
}

func I642S(i64 int64) string {
	s := strconv.FormatInt(i64, 10)
	return s
}

func I322S(i32 int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(i32)
	signed := i < 0
	if signed {
		i = -i
	}

	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func SSlice2I64Slice(ss []string) []int64 {
	if ss == nil {
		return nil
	}
	ii := make([]int64, 0, len(ss))
	for _, s := range ss {
		ii = append(ii, S2I64(s))
	}
	return ii
}

func I64Slice2SSlice(ii []int64) []string {
	if ii == nil {
		return nil
	}
	ss := make([]string, 0, len(ii))
	for _, i := range ii {
		ss = append(ss, I642S(i))
	}
	return ss
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return ""
}
