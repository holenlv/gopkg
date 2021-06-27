package concurrent

import (
	"log"
	"runtime/debug"
)

func SafeGo(fn func()) {
	go func() {
		defer Recovery()
		fn()
	}()
}
func Recovery() {
	if err := recover(); err != nil {
		log.Printf("catch panic,stack trace=%v", string(debug.Stack()))
	}
	return
}
