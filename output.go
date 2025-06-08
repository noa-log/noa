/*
 * @Author: nijineko
 * @Date: 2025-06-08 10:59:01
 * @LastEditTime: 2025-06-08 15:47:09
 * @LastEditors: nijineko
 * @Description: log output package
 * @FilePath: \noa\output.go
 */
package noa

import (
	"fmt"
	"time"

	"github.com/noa-log/colorize"
	"github.com/noa-log/noa/errors"
)

/**
 * @description: print debug log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Debug(Source string, Data ...any) {
	l.Println(DEBUG, Source, Data...)
}

/**
 * @description: print info log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Info(Source string, Data ...any) {
	l.Println(INFO, Source, Data...)
}

/**
 * @description: print warning log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Warning(Source string, Data ...any) {
	l.Println(WARNING, Source, Data...)
}

/**
 * @description: print error log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Error(Source string, Data ...any) {
	for Index, Value := range Data {
		// If the type is error, then wrap it
		if Error, ok := Value.(error); ok {
			WrapError := errors.Wrap(Error, l.Errors.CallerSkip)
			Data[Index] = WrapError
		}
	}

	l.Println(ERROR, Source, Data...)
}

/**
 * @description: print fatal log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Fatal(Source string, Data ...any) {
	for Index, Value := range Data {
		// If the type is error, then wrap it
		if Error, ok := Value.(error); ok {
			WrapError := errors.Wrap(Error, l.Errors.CallerSkip)
			Data[Index] = WrapError
		}
	}

	l.Println(FATAL, Source, Data...)
}

/**
 * @description: Print log with newline
 * @param {int} Level Log level
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Println(Level int, Source string, Data ...any) {
	// Check if the last element is a string and ends with a newline
	if l.AutoLastNewline && len(Data) > 0 {
		if StrValue, ok := Data[len(Data)-1].(string); ok && len(StrValue) > 0 && StrValue[len(StrValue)-1] == '\n' {
			l.Print(Level, Source, Data...)
			return
		}
	}

	Data = append(Data, "\n")
	l.Print(Level, Source, Data...)
}

/**
 * @description: Print log
 * @param {int} Level Log level
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Print(Level int, Source string, Data ...any) {
	if Level < l.Level || Level == OFF {
		return
	}

	// Execute before handles
	for _, Handle := range l.beforeHandles {
		// Convert []any to []*any
		DataPtrs := make([]*any, len(Data))
		for i := range Data {
			DataPtrs[i] = &Data[i]
		}
		if err := Handle(&Level, &Source, DataPtrs...); err != nil {
			panic(err)
		}
	}

	var PrintData []any

	// Add time
	PrintData = append(PrintData, colorize.CyanText(
		time.Now().Format(l.TimeFormat),
	))

	// Add log level
	switch Level {
	case DEBUG:
		PrintData = append(PrintData, colorize.GrayText("DEBUG"))
	case INFO:
		PrintData = append(PrintData, colorize.BlueText("INFO"))
	case WARNING:
		PrintData = append(PrintData, colorize.YellowText("WARNING"))
	case ERROR:
		PrintData = append(PrintData, colorize.RedText("ERROR"))
	case FATAL:
		PrintData = append(PrintData, colorize.MagentaText("FATAL"))
	}

	// Add source
	PrintData = append(PrintData, "["+Source+"]")

	// Add data
	PrintData = append(PrintData, Data...)

	// Append the stack trace when wrapping the errors
	if l.Errors.StackTrace {
		ErrorStackData := make([]any, 0, len(PrintData))
		for _, Value := range PrintData {
			if WrapError, ok := Value.(*errors.Error); ok {
				// Append the stack trace to the data
				ErrorStackData = append(ErrorStackData, Value, "\n", WrapError.StackFormat())
			} else {
				// Append the original value
				ErrorStackData = append(ErrorStackData, Value)
			}
		}

		PrintData = ErrorStackData
	}

	// Pad spaces in the data slice
	PrintData = padSpace(PrintData)

	// Remove color data
	RemoveColorData := make([]any, len(PrintData))
	for Index, Value := range PrintData {
		if StrValue, ok := Value.(string); ok {
			RemoveColorData[Index] = colorize.Remove(StrValue)
		} else {
			RemoveColorData[Index] = Value
		}
	}
	// if colorRemove is true, use the data without color
	if l.RemoveColor {
		PrintData = RemoveColorData
	}

	// Print log
	fmt.Print(PrintData...)

	// Write to file if enabled
	if l.Writer.Enable {
		LogFileHandle, err := l.Writer.openFile()
		if err != nil {
			return
		}

		// Write log to file
		if _, err := fmt.Fprint(LogFileHandle, RemoveColorData...); err != nil {
			return
		}
	}

	// Execute after handles
	for _, Handle := range l.bfterHandles {
		if err := Handle(Level, Source, Data...); err != nil {
			panic(err)
		}
	}
}

/**
 * @description: Pad spaces in the data slice
 * @param {[]any} Data Data slice to pad spaces in
 * @return {[]any} Padded data slice
 */
func padSpace(Data []any) []any {
	ResultData := make([]any, 0, len(Data)*2-1)
	for Index, Value := range Data {
		ResultData = append(ResultData, Value)
		if Index < len(Data)-1 {
			// If the value is a string and ends with a newline, skip adding a space
			if StrValue, ok := Value.(string); ok && (len(StrValue) > 0 && StrValue[len(StrValue)-1] == '\n') {
				continue
			}

			ResultData = append(ResultData, " ")
		}
	}
	return ResultData
}
