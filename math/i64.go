package math

func MaxI64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MinI64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func AbsI64(a int64) int64 {
	if a > 0 {
		return a
	}
	return a * (-1)
}
