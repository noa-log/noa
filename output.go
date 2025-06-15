/*
 * @Author: nijineko
 * @Date: 2025-06-08 10:59:01
 * @LastEditTime: 2025-06-15 10:49:36
 * @LastEditors: nijineko
 * @Description: log output package
 * @FilePath: \noa\output.go
 */
package noa

import (
	"time"

	"github.com/noa-log/noa/encoder"
	"github.com/noa-log/noa/errors"
)

/**
 * @description: print debug log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Debug(Source string, Data ...any) {
	l.Print(DEBUG, Source, Data...)
}

/**
 * @description: print info log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Info(Source string, Data ...any) {
	l.Print(INFO, Source, Data...)
}

/**
 * @description: print warning log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Warning(Source string, Data ...any) {
	l.Print(WARNING, Source, Data...)
}

/**
 * @description: print error log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Error(Source string, Data ...any) {
	l.Print(ERROR, Source, Data...)
}

/**
 * @description: print fatal log
 * @param {string} Source Log source (e.g., file name, function name)
 * @param {...any} Data print data
 */
func (l *LogConfig) Fatal(Source string, Data ...any) {
	l.Print(FATAL, Source, Data...)
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

	// wrap errors
	if Level == ERROR || Level == FATAL {
		for Index, Value := range Data {
			// If the type is error, then wrap it
			if Error, ok := Value.(error); ok {
				WrapError := errors.Wrap(Error, l.Errors.CallerSkip)
				Data[Index] = WrapError
			}
		}
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

	// Create encoder context
	EncoderContext := encoder.NewContext(time.Now(), Level, Source, Data)

	// print log data
	l.Encoder.Print.Print(EncoderContext)

	// Write to file if enabled
	if l.Writer.Enable {
		LogFileHandle, err := l.Writer.openFile(l.Encoder.Write.WriteFileExtension())
		if err != nil {
			panic(err)
		}

		// Write log to file
		if err := l.Encoder.Write.Write(LogFileHandle, EncoderContext); err != nil {
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
