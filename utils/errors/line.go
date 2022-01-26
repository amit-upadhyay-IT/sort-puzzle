package errors

import (
	"fmt"
	"runtime"
)

type lineError struct {
	line int
	file string
	err  error
}

func (e *lineError) Error() string {
	for i := len(e.file) - 1; i > 0; i-- {
		if e.file[i] == '/' {
			return fmt.Sprintf("%s:%d: %s", e.file[i+1:], e.line, e.err.Error())
		}
	}
	return e.err.Error()
}

func Errorf(msg string, a ...interface{}) error {
	err := fmt.Errorf(msg, a)
	_, file, line, _ := runtime.Caller(1)
	return &lineError{
		line: line,
		file: file,
		err:  err,
	}
}
