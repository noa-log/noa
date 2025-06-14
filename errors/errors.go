/*
 * @Author: nijineko
 * @Date: 2025-06-08 12:42:57
 * @LastEditTime: 2025-06-12 11:10:17
 * @LastEditors: nijineko
 * @Description: noa errors package
 * @FilePath: \noa\errors\errors.go
 */
package errors

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	MAX_STACK_DEPTH = 64
)

// Error structure
type Error struct {
	Err   error     // base error
	Stack []uintptr // stack trace

	stackFrames []StackFrame // stack frames
}

/**
 * @description: Create a new error
 * @param {string} Text error message
 * @param {int} Skip number of stack frames to skip (default is 1)
 * @return {*Error} wrapped error
 */
func New(Text string, Skip int) *Error {
	return Wrap(errors.New(Text), Skip+1)
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
 * @description: Check if the error is of a specific type
 * @param {Error} Err error to check
 * @param {Error} Target target error
 * @return {bool} true if Err is of type Target, false otherwise
 */
func Is(Err *Error, Target *Error) bool {
	if Err.Err == nil {
		return false
	}
	if Target.Err == nil {
		return false
	}

	return errors.Is(Err.Err, Target.Err)
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
func (e Error) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

/**
 * @description: Get the stack trace as a formatted string
 * @return {string} formatted stack trace
 */
func (e *Error) StackFormat() string {
	if e.Err == nil {
		return ""
	}

	Frames := e.StackFrames()
	var StackString string
	for _, Frame := range Frames {
		StackString += fmt.Sprintf("%s\n\t%s:%d +0x%x\n",
			Frame.Function,
			Frame.File,
			Frame.Line,
			Frame.PC-Frame.Entry,
		)
	}

	return StackString
}
