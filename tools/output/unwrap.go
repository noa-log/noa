/*
 * @Author: nijineko
 * @Date: 2025-06-10 12:44:11
 * @LastEditTime: 2025-06-10 13:22:42
 * @LastEditors: nijineko
 * @Description: output print data unwrap utility
 * @FilePath: \noa\tools\output\unwrap.go
 */
package output

import (
	"time"

	"github.com/noa-log/noa"
)

/**
 * @description: Unwrap print data from the log output.
 * @param {noa.LogConfig} Log noa log instance
 * @param {[]any} PrintData data to unwrap
 * @return {time.Time} parsed log time
 * @return {int} parsed log level
 * @return {string} parsed log source
 * @return {[]any} unwrapped print data (remove newline at the end)
 */
func UnwrapPrintData(Log *noa.LogConfig, PrintData []any) (time.Time, int, string, []any) {
	var (
		Time       time.Time
		Level      int
		Source     string
		PrintArray []any
	)

	if len(PrintData) < 3 {
		return Time, Level, Source, PrintArray
	}

	// parse log time
	if TimeStr, ok := PrintData[0].(string); ok {
		TimeData, err := time.ParseInLocation(Log.TimeFormat, TimeStr, time.Local)
		if err == nil {
			Time = TimeData
		}
	}

	// parse log level
	if LevelStr, ok := PrintData[1].(string); ok {
		switch LevelStr {
		case "DEBUG":
			Level = noa.DEBUG
		case "INFO":
			Level = noa.INFO
		case "WARNING":
			Level = noa.WARNING
		case "ERROR":
			Level = noa.ERROR
		case "FATAL":
			Level = noa.FATAL
		default:
			Level = -1 // Unknown level
		}
	}

	// parse log source
	if SourceStr, ok := PrintData[2].(string); ok {
		if len(SourceStr) >= 2 && SourceStr[0] == '[' && SourceStr[len(SourceStr)-1] == ']' {
			Source = SourceStr[1 : len(SourceStr)-1]
		} else {
			Source = SourceStr
		}
	}

	// remove the last line break
	if PrintData[len(PrintData)-1] == "\n" {
		PrintData = PrintData[:len(PrintData)-1]
	}

	// parse print data
	if len(PrintData) > 3 {
		PrintArray = PrintData[3:]
	}

	return Time, Level, Source, PrintArray
}
