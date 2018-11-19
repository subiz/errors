package errors

import (
	"testing"
)

type mes struct {
	s string
}

func (m mes) String() string { return m.s }

func TestMarshal(t *testing.T) {
	err := New(500, E_unknown, "thanh %d", 4)
	derr := FromString(err.Error())
	if err.Class != derr.Class ||
		err.Stack != derr.Stack ||
		err.Created != derr.Created ||
		err.Code != derr.Code ||
		err.RequestId != derr.RequestId {
		t.Errorf("wrong")
	}
}

func foo() error   { return bar() }
func bar() error  { return New(500, E_unknown) }
func hello() error { return foo() }

func TestErrorStack(t *testing.T) {
	//trace()
	err := hello().(*Error)
	println(err.Stack)
	//e := New(400, &mes{"hi"})
	//println(e.Error())
}

func TestIsSystemPath(t *testing.T) {
	tcs := []struct {
		path   string
		expect bool
	}{
		{"/usr/local/go/src/runtime/debug/stack.go:24", true},
		{"", false},
		{"/root/go/src/git.subiz.net/api/vendor/errors/errors.go:133", false},
	}

	for _, tc := range tcs {
		if out := isSystemPath(tc.path); out != tc.expect {
			t.Errorf("expect %v, got %v for path %s", tc.expect, out, tc.path)
		}
	}
}

func TestTrimOutPrefix(t *testing.T) {
	tcs := []struct {
		str, prefix, expect string
	}{
		{"123b456", "b", "456"},
		{"123  456", " ", " 456"},
		{"12b3b456", "b", "3b456"},
		{"", "b", ""},
		{"123", "b", "123"},
		{"123", "123", ""},
		{"123", "1234", "123"},
		{"123", "23", ""},
		{"123", "5", "123"},
	}

	for _, tc := range tcs {
		if out := trimOutPrefix(tc.str, tc.prefix); out != tc.expect {
			t.Errorf("expect %v, got %v, for str %s and prefix %s", tc.expect, out, tc.str, tc.prefix)
		}
	}
}

func TestTrimToPrefix(t *testing.T) {
	tcs := []struct {
		str, prefix, expect string
	}{
		{"123b456", "b", "b456"},
		{"123  456", " ", "  456"},
		{"12b3b456", "b", "b3b456"},
		{"", "b", ""},
		{"123", "b", "123"},
		{"123", "123", "123"},
		{"123", "1234", "123"},
		{"123", "23", "23"},
	}

	for _, tc := range tcs {
		if out := trimToPrefix(tc.str, tc.prefix); out != tc.expect {
			t.Errorf("expect %v, got %v, for str %s and prefix %s", tc.expect, out, tc.str, tc.prefix)
		}
	}
}

func BenchmarkGetStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getStack(0)
	}
}


func BenchmarkTrimToPrefix(b *testing.B) {
	for n := 0; n < b.N; n++ {
		trimToPrefix("12b3b456", "b")
	}
}

func BenchmarkIsSystemPath(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isSystemPath("/usr/local/go/src")
		isSystemPath("abcdef/git.subiz.net/errors/")
	}
}
