package errors

import (
	"encoding/json"
	"hash/crc32"
	"runtime/debug"
	"strings"
	"time"
)

type stringer struct {
	s string
}

func (s stringer) String() string { return s.s }

var crc32q = crc32.MakeTable(0xD5828281)

func Default(e interface{}) *Error {
	if e == nil {
		return nil
	}

	e1, ok := e.(*Error)
	if ok {
		return e1
	}

	e2, ok := e.(error)
	if ok {
		return Wrap(e2, 500, stringer{"unknown"})
	}

	return New(500, stringer{"unknown"}, Sprintf("%v", e))
}

func Wrap(err error, class int, code Stringer, v ...interface{}) *Error {
	if err == nil {
		err = &Error{}
	}
	mye, ok := err.(*Error)
	if !ok {
		e := New(class, code, append(v, err.Error()))
		e.Base = New(500, stringer{"base"}, err)
		return e
	}

	if code.String() != "" && (mye.Code == "" || mye.Code == "unknown") {
		mye.Code = code.String()
	}

	if class != 0 && mye.Class == 0 {
		mye.Class = int32(class)
	}

	if len(v) > 0 {
		e := New(class, code, v)
		mye.Description += "\n" + e.Description
	}
	return mye
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(class int, code Stringer, v ...interface{}) *Error {
	var format, message string
	if len(v) == 0 {
		format = ""
	} else {
		var ok bool
		format, ok = v[0].(string)
		if !ok {
			format = strings.Repeat("%v", len(v))
		} else {
			v = v[1:]
		}
	}
	message = Sprintf(format, v...)

	stack := getStack(5)
	e := &Error{}
	e.Description = message
	e.Class = int32(class)
	e.Stack = string(stack)
	e.Created = time.Now().UnixNano()
	e.Code = code.String()
	e.Hash = Sprintf("%08x", crc32.Checksum(stack, crc32q))
	e.Base = &Error{}
	return e
}

func FromError(err string) *Error {
	if !strings.HasPrefix(err, "#ERR ") {
		return New(500, stringer{"unknown"}, err)
	}
	e := &Error{}
	if er := json.Unmarshal([]byte(err[len("#ERR "):]), e); er != nil {
		return New(500, stringer{"invalid_json_err"}, "%s, %s", er, err)
	}
	return e
}

func (e *Error) GetCode() string {
	if e == nil {
		return ""
	}

	return e.Code
}

func (e *Error) Interface() error {
	if e == nil {
		return nil
	}
	return e
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	b, err := json.Marshal(e)
	if err != nil {
		return "#ERRX " + err.Error() + "(" + strings.Replace(string(getStack(5)), "\n", "|", -1) + ")"
	}
	return "#ERR " + string(b)
}

func getStack(skip int) []byte {
	s := string(debug.Stack())
	lines := strings.Split(strings.TrimSpace(s), "\n")
	if len(lines) <= skip {
		skip = len(lines)
	}
	lines = lines[skip:] // ignore unnecessary lines
	out := ""
	for i, line := range lines {
		if i%2 == 1 { // filter lines contains file path
			f := removeLastPlusSign(strings.TrimSpace(line))
			f = splitLineNumber(f)
			out += f + "\n"
		}
	}
	return []byte(out)
}

func removeLastPlusSign(s string) string {
	split := strings.Split(s, " ")
	if len(split) < 2 {
		return s
	}
	if !strings.HasPrefix(split[len(split)-1], "+0x") {
		return s
	}
	return strings.Join(split[0:len(split)-1], " ")
}

func splitLineNumber(s string) string {
	split := strings.Split(s, ":")
	if len(split) < 2 {
		return s
	}

	line := split[len(split)-1]
	return strings.Join(split[0:len(split)-1], ":") + ":" + line
}
