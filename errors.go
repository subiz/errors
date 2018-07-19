package errors

import (
	"encoding/json"
	"hash/crc32"
	"runtime/debug"
	"strings"
	"time"
)

var crc32q = crc32.MakeTable(0xD5828281)

func Wrap(err error, class int, code Stringer, v ...interface{}) *Error {
	mye, ok := err.(*Error)
	if !ok {
		e := New(class, code, append(v, err.Error()))
		e.Base = err
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

	stack := getStack()
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

func (e *Error) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "#ERRX " + err.Error() + "(" + strings.Replace(string(getStack()), "\n", "|", -1) + ")"
	}
	return "#ERR " + string(b)
}

func getStack() []byte {
	s := string(debug.Stack())
	lines := strings.Split(strings.TrimSpace(s), "\n")
	lines = lines[3:] // ignore unnecessary lines
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
