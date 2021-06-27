package concurrent

import "testing"

func TestSafeGo(t *testing.T) {
	SafeGo(func() {
		panic("hello")
	})
	select {}
}
