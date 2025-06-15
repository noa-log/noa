/*
 * @Author: nijineko
 * @Date: 2025-06-08 10:29:01
 * @LastEditTime: 2025-06-15 10:56:23
 * @LastEditors: nijineko
 * @Description: noa log package
 * @FilePath: \noa\log.go
 */
package noa

import (
	"os"
	"time"

	"github.com/noa-log/noa/encoder"
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
	Enable     bool   // enable log file writing
	FolderPath string // folder path for log files
	TimeFormat string // time format for log file names

	file map[string]*os.File // log file handles
}

// Log config encoder structure
type LogConfigEncoder struct {
	Print encoder.Encoder
	Write encoder.Encoder
}

// Log config structure
type LogConfig struct {
	Level       int              // log level
	RemoveColor bool             // remove color from log output
	TimeFormat  string           // log prefix time format
	Errors      LogConfigErrors  // error configuration for logging
	Writer      LogConfigWriter  // writer configuration for logging to files
	Encoder     LogConfigEncoder // log encoder

	beforeHandles []BeforeHandleFunc // functions to run before logging
	bfterHandles  []AfterHandleFunc  // functions to run after logging
}

/**
 * @description: Create a new log configuration instance
 * @return {*LogConfig} a log configuration instance
 */
func NewLog() *LogConfig {
	Config := &LogConfig{
		Level:       DEBUG,
		RemoveColor: false,
		TimeFormat:  "2006-01-02 15:04:05",
		Errors: LogConfigErrors{
			StackTrace: true,
			CallerSkip: 4, // default skip 4 frames to find the caller
		},
		Writer: LogConfigWriter{
			Enable:     true,
			FolderPath: "./logs",
			TimeFormat: "2006-01-02",
			file:       make(map[string]*os.File), // initialize file map
		},
	}
	Config.SetEncoder(NewTextEncoder(Config)) // default text encoder

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

/**
 * @description: Set encoder for the log instance
 * @param {encoder.Encoder} Encoder encoder instance
 */
func (l *LogConfig) SetEncoder(Encoder encoder.Encoder) {
	l.Encoder.Print = Encoder
	l.Encoder.Write = Encoder
}
