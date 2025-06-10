/*
 * @Author: nijineko
 * @Date: 2025-06-10 11:30:17
 * @LastEditTime: 2025-06-10 11:54:47
 * @LastEditors: nijineko
 * @Description: error stack frame
 * @FilePath: \noa\errors\stackFrame.go
 */
package errors

import (
	"runtime"
	"strings"
)

// Stack Frame structure
type StackFrame struct {
	PC           uintptr // Program Counter
	Function     string  // Function name
	File         string  // File name
	Line         int     // Line number
	Entry        uintptr // Entry point of the function
	PackageName  string  // Package name
	FunctionName string  // Function name without package
}

/**
 * @description: Get the stack frames of the error
 * @return {[]StackFrame} stack frames
 */
func (e *Error) StackFrames() []StackFrame {
	if e.stackFrames != nil {
		return e.stackFrames
	}

	CallersFrames := runtime.CallersFrames(e.Stack)
	var StackFrames []StackFrame

	for {
		Frame, More := CallersFrames.Next()
		if Frame.Function == "" {
			break
		}

		// Parse the function to get package name and function name
		PackageName, FunctionName := ParseFunction(Frame.Func)

		StackFrames = append(StackFrames, StackFrame{
			PC:           Frame.PC,
			Function:     Frame.Function,
			File:         Frame.File,
			Line:         Frame.Line,
			Entry:        Frame.Entry,
			PackageName:  PackageName,
			FunctionName: FunctionName,
		})
		if !More {
			break
		}
	}

	e.stackFrames = StackFrames
	return StackFrames
}

/**
 * @description: Parse runtime.Func to get the package name and function name
 * @param {*runtime.Func} Func runtime function
 * @return {string} package name
 * @return {string} function name
 */
func ParseFunction(Func *runtime.Func) (string, string) {
	FuncName := Func.Name()

	var PackageName, FunctionName string
	FuncNameArr := strings.Split(FuncName, ".")
	if len(FuncNameArr) > 1 {
		FunctionName = FuncNameArr[len(FuncNameArr)-1]
		PackageName = strings.Join(FuncNameArr[:len(FuncNameArr)-1], ".")
	} else {
		// if the function name does not contain a package name
		FunctionName = FuncName
	}

	return PackageName, FunctionName
}
