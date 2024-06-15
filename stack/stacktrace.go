package stack

import (
	"fmt"
	"runtime"
)

type ErrStackTrace struct {
	cause error
	callers []caller
}

func (e ErrStackTrace) Error() string {
	msg := fmt.Sprintf("[ERROR] %s\n", e.cause.Error())
	for _, c := range e.callers {
		msg += c.String()
	}
	return msg
}

func (e ErrStackTrace) Unwrap() error {
	return e.cause
}

type caller struct {
	file string
	line int
	function string
}

func (c caller) String() string {
	return fmt.Sprintf("file: %v, line: %v, func %v()\n", c.file, c.line, c.function)
}

// errorを引数として受けとりラップして、コールスタックを出力できるerror型であるErrCallStackを返却する
func NewCallStack(cause error) error {
	pcs := make([]uintptr, 32)
	num := runtime.Callers(2, pcs)
	callers := make([]caller, num)
	for i := 0; i < num; i++ {
		fun := runtime.FuncForPC(pcs[i])
		file, line := fun.FileLine(pcs[i])
		function := fun.Name()
		callers[i] = caller{
			file: file,
			line: line,
			function: function,
		}
	}
	return ErrStackTrace{
		cause: cause,
		callers: callers,
	}
}