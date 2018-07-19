package errors

import (
	"runtime"
	_ "strings"
	"testing"
)

type mes struct {
	s string
}

func (m mes) String() string { return m.s }

func TestErrorStack(t *testing.T) {
	trace()
	e := New(400, &mes{"hi"})
	println(e.Error())
}

func TestErrorWrapStack(t *testing.T) {
	trace()
	err := Errorf("thanh %d", 4)
	e := Wrap(err, 400, &mes{"hi"}, "a", "b")
	println(e.Error())
}

func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	Printf("\nTEST %s\n", f.Name())
}

func TestErrorWrapError(t *testing.T) {
	trace()
	err := Errorf("thanh %d", 4)
	e := Wrap(err, 0, &mes{"hi"}, "a", "b")
	e2 := Wrap(e, 500, &mes{"hi2"}, "c", "d")
	println(e2.Error())
}
