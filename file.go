/*
 * @Author: nijineko
 * @Date: 2025-06-08 11:03:48
 * @LastEditTime: 2025-06-08 12:38:05
 * @LastEditors: nijineko
 * @Description: log file handle utility package
 * @FilePath: \noa\file.go
 */
package noa

import (
	"os"
	"path/filepath"
	"time"
)

/**
 * @description: Open a log file handle
 * @return {*os.File} log file handle
 * @return {error} error
 */
func (lcw *LogConfigWriter) openFile() (*os.File, error) {
	FileNameTime := time.Now().Format(lcw.TimeFormat)
	// check file handle already exists
	if _, ok := lcw.file[FileNameTime]; ok {
		return lcw.file[FileNameTime], nil
	}

	// create new file handle
	FileName := FileNameTime + ".log"
	FilePath := filepath.Join(lcw.FolderPath, FileName)
	// create folder if it doesn't exist
	if _, err := os.Stat(lcw.FolderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(lcw.FolderPath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	LogFileHandle, err := os.OpenFile(FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	lcw.file[FileNameTime] = LogFileHandle
	return LogFileHandle, nil
}

/**
 * @description: Close unused log files
 */
func (lcw *LogConfigWriter) closeUnusedFiles() {
	for fileName, fileHandle := range lcw.file {
		Timestamp, err := time.Parse(lcw.TimeFormat, fileName)
		if err != nil {
			continue
		}

		// Close file if it's not from today
		if !isToday(Timestamp) {
			fileHandle.Close()
			// Remove file handle from map
			delete(lcw.file, fileName)
		}
	}
}

/**
 * @description: Check if the given time is today
 * @param {time.Time} Time Time to check
 * @return {bool} true if the time is today, false otherwise
 */
func isToday(Time time.Time) bool {
	YNow, MNow, DNow := time.Now().Date()
	YTime, MTime, DTime := Time.Date()
	return YNow == YTime && MNow == MTime && DNow == DTime
}
