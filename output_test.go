/*
 * @Author: nijineko
 * @Date: 2025-06-08 13:16:22
 * @LastEditTime: 2025-06-08 15:00:38
 * @LastEditors: nijineko
 * @Description: log output test
 * @FilePath: \noa\output_test.go
 */
package noa

import (
	"errors"
	"testing"
)

func TestPrintLog(t *testing.T) {
	Log := NewLog()

	Log.Print(DEBUG, "Test", "This is a debug log")
	Log.Print(DEBUG, "Test", "This is a debug log")
	Log.Print(INFO, "Test", "This is an info log")
	Log.Print(WARNING, "Test", "This is a warning log")
	Log.Print(ERROR, "Test", "This is an error log")
	Log.Print(FATAL, "Test", "This is a fatal log")
}

func TestPrintLevelLog(t *testing.T) {
	Log := NewLog()

	Log.Debug("Test", "This is a debug log")
	Log.Info("Test", "This is an info log")
	Log.Warning("Test", "This is a warning log")
	Log.Error("Test", "This is an error log")
	Log.Fatal("Test", "This is a fatal log")
}

func TestPrintErrorLog(t *testing.T) {
	Log := NewLog()
	Log.Errors.CallerSkip = 3

	// Test with a simple error
	err := errors.New("This is a test error")
	Log.Error("Test", err, "Test append")

	// Test with multiple errors
	err1 := errors.New("This is the first test error")
	err2 := errors.New("This is the second test error")
	Log.Error("Test", err1, err2)
}

func TestStress(t *testing.T) {
	Log := NewLog()

	for i := 0; i < 10000; i++ {
		Log.Debug("Test", "This is a debug log", i)
		Log.Info("Test", "This is an info log", i)
		Log.Warning("Test", "This is a warning log", i)
		Log.Error("Test", "This is an error log", i)
		Log.Fatal("Test", "This is a fatal log", i)
	}
}
