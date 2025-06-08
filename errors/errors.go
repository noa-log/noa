/*
 * @Author: nijineko
 * @Date: 2025-06-08 12:42:57
 * @LastEditTime: 2025-06-08 14:37:04
 * @LastEditors: nijineko
 * @Description: noa errors package
 * @FilePath: \noa\errors\errors.go
 */
package errors

import (
	"fmt"
	"runtime"
)

const (
	MAX_STACK_DEPTH = 64
)

// Error structure
type Error struct {
	Err   error     // base error
	Stack []uintptr // stack trace
}

func (e *Error) StackString() {
	panic("unimplemented")
}

/**
 * @description: Wrap an error with a stack trace
 * @param {error} Err base error
 * @param {int} Skip number of stack frames to skip
 * @return {*Error} wrapped error with stack trace
 */
func Wrap(Err error, Skip int) *Error {
	Stack := make([]uintptr, MAX_STACK_DEPTH)
	Length := runtime.Callers(Skip, Stack[:])
	return &Error{
		Err:   Err,
		Stack: Stack[:Length],
	}
}

/**
 * @description: Unwrap the error to get the original error
 * @return {error} original error
 */
func (e *Error) Unwrap() error {
	return e.Err
}

/**
 * @description: Get error message
 * @return {string} error message
 */
func (e *Error) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

/**
 * @description: Get the stack trace of the error
 * @return {[]uintptr} stack trace
 */
func (e *Error) StackTrace() []uintptr {
	return e.Stack
}

/**
 * @description: Get the stack trace as a formatted string
 * @return {string} formatted stack trace
 */
func (e *Error) StackFormat() string {
	if e.Err == nil {
		return ""
	}

	Frames := runtime.CallersFrames(e.Stack)
	var StackString string
	for {
		Frame, More := Frames.Next()
		if Frame.Function == "" {
			break
		}
		StackString += fmt.Sprintf("%s\n\t%s:%d +0x%x\n",
			Frame.Function,
			Frame.File,
			Frame.Line,
			Frame.PC-Frame.Entry,
		)
		if !More {
			break
		}
	}

	return StackString
}
