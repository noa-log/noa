/*
 * @Author: nijineko
 * @Date: 2025-06-08 10:29:01
 * @LastEditTime: 2025-06-08 13:43:49
 * @LastEditors: nijineko
 * @Description: noa log package
 * @FilePath: \noa\log.go
 */
package noa

import (
	"os"
	"time"

	"github.com/noa-log/noa/file/encoder"
	"github.com/noa-log/noa/file/encoder/text"
)

// default log levels
const (
	DEBUG = iota + 0
	INFO
	WARNING
	ERROR
	FATAL
	OFF
)

// Log config error structure
type LogConfigErrors struct {
	StackTrace bool // whether to print stack trace for errors
	CallerSkip int  // skip number of stack frames to find the caller
}

// Log config writer structure
type LogConfigWriter struct {
	Enable     bool            // enable log file writing
	FolderPath string          // folder path for log files
	TimeFormat string          // time format for log file names
	Encoder    encoder.Encoder // encoder for log file writing

	file map[string]*os.File // log file handles
}

// Log config structure
type LogConfig struct {
	Level           int             // log level
	RemoveColor     bool            // remove color from log output
	TimeFormat      string          // log prefix time format
	AutoLastNewline bool            // auto Check if the last element ends with a newline and skip appending if present
	Errors          LogConfigErrors // error configuration for logging
	Writer          LogConfigWriter // writer configuration for logging to files

	beforeHandles []BeforeHandleFunc // functions to run before logging
	bfterHandles  []AfterHandleFunc  // functions to run after logging
}

/**
 * @description: Create a new log configuration instance
 * @return {*LogConfig} a log configuration instance
 */
func NewLog() *LogConfig {
	Config := &LogConfig{
		Level:           DEBUG,
		RemoveColor:     false,
		TimeFormat:      "2006-01-02 15:04:05",
		AutoLastNewline: true,
		Errors: LogConfigErrors{
			StackTrace: true,
			CallerSkip: 3, // default skip 3 frames to find the caller
		},
		Writer: LogConfigWriter{
			Enable:     true,
			FolderPath: "./logs",
			TimeFormat: "2006-01-02",
			Encoder:    text.NewTextEncoder(),     // default encoder is TextEncoder
			file:       make(map[string]*os.File), // initialize file map
		},
	}

	// Enabled a goroutine to periodically close unused log file handles
	go func() {
		for {
			// Close unused log files every minute
			Config.Writer.closeUnusedFiles()

			time.Sleep(time.Minute)
		}
	}()

	return Config
}
