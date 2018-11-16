package errors

import (
	"encoding/json"
	"hash/crc32"
	F "fmt"
	// "reflect"
	"runtime/debug"
	"strings"
	"time"
)

var crc32q = crc32.MakeTable(0xD5828281)

// Default converts e to *Error
/*func Default(e interface{}) *Error {
	e1, ok := e.(*Error)
	if ok {
		return e1
	}

	err, ok := e.(error)
	if ok {
		return Wrap(err, 500, E_unknown)
	}

	if e == nil || reflect.ValueOf(e).IsNil() {
		return nil
	}

	return New(500, E_unknown, Sprintf("%v", e))
}*/

func Wrap(err error, class int, code Stringer, v ...interface{}) *Error {
	if err == nil {
		err = &Error{}
	}
	mye, ok := err.(*Error)
	if !ok {
		e := New(class, code, append(v, err.Error()))
		e.Root = err.Error()
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

	stack := getStack(7)
	e := &Error{}
	e.Description = message
	e.Class = int32(class)
	e.Stack = string(stack)
	e.Created = time.Now().UnixNano()
	e.Code = code.String()
	e.Hash = F.Sprintf("%08x", crc32.Checksum(stack, crc32q))
	return e
}

// FromString unmarshal an error string to *Error
func FromString(err string) *Error {
	if !strings.HasPrefix(err, "#ERR ") {
		return New(500, E_unknown, err)
	}
	e := &Error{}
	if er := json.Unmarshal([]byte(err[len("#ERR "):]), e); er != nil {
		return New(500, E_json_marshal_error, "%s, %s", er, err)
	}
	return e
}

// GetCode returns code of the error
func (e *Error) GetCode() string {
	if e == nil {
		return ""
	}

	return e.Code
}

// Interface returns error interface of *Error.
// If e is nil return interface(nil, nil) instead of interface(*Error, nil) so
// the check `if e.Interface() == nil {}` will be true
func (e *Error) Interface() error {
	if e == nil {
		return nil
	}
	return e
}

// Error returns string representation of an Error
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
