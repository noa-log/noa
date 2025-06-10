/*
 * @Author: nijineko
 * @Date: 2025-06-10 21:24:26
 * @LastEditTime: 2025-06-10 22:29:29
 * @LastEditors: nijineko
 * @Description: text log encoder
 * @FilePath: \noa\textEncoder.go
 */
package noa

import (
	"fmt"
	"os"

	"github.com/noa-log/colorize"
	"github.com/noa-log/noa/encoder"
	"github.com/noa-log/noa/errors"
	"github.com/noa-log/noa/tools"
)

type TextEncoderConfigNewline struct {
	Auto  bool // enable automatic newline at the end of each log entry
	Smart bool // enable smart newline handling, auto checks if the last element ends with a newline and skips appending if present
}

// TextEncoder struct
type TextEncoder struct {
	Newline TextEncoderConfigNewline // newline configuration

	Log *LogConfig // noa log instance
}

/**
 * @description: Create a new TextEncoder instance
 * @return {*TextEncoder} TextEncoder instance
 */
func NewTextEncoder(Log *LogConfig) *TextEncoder {
	return &TextEncoder{
		Log: Log,
		Newline: TextEncoderConfigNewline{
			Auto:  true,
			Smart: true,
		},
	}
}

// print log data
func (te *TextEncoder) Print(c *encoder.Context) {
	// generate text print data
	PrintData, RemoveColorData := generateTextPrintData(te.Log, c, te.Newline)

	// Save removed color data to context
	c.Set("PrintData", RemoveColorData)

	// Print log
	fmt.Print(PrintData...)
}

// return file extension for the encoded file
func (te *TextEncoder) WriteFileExtension() string {
	return ".log"
}

// write log data to file
func (te *TextEncoder) Write(FileHandle *os.File, c *encoder.Context) error {
	RemoveColorData := c.Get("PrintData")
	RemoveColorSlice, ok := RemoveColorData.([]any)
	if RemoveColorData == nil || !ok {
		_, RemoveColorSlice = generateTextPrintData(te.Log, c, te.Newline)
	}

	if _, err := fmt.Fprint(FileHandle, RemoveColorSlice...); err != nil {
		return err
	}

	return nil
}

/**
 * @description: Generate text print data for logging
 * @param {*LogConfig} Log noa log instance
 * @param {*encoder.Context} c log encoder context
 * @param {TextEncoderConfigNewline} NewlineConfig newline configuration
 * @return {[]any} PrintData log data to print
 * @return {[]any} RemoveColorData log data without color
 */
func generateTextPrintData(Log *LogConfig, c *encoder.Context, NewlineConfig TextEncoderConfigNewline) ([]any, []any) {
	var PrintData []any

	// Add time
	PrintData = append(PrintData, colorize.CyanText(
		c.Time.Format(Log.TimeFormat),
	))

	// Add log level
	switch c.Level {
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
	PrintData = append(PrintData, "["+c.Source+"]")

	// Add data
	PrintData = append(PrintData, c.Data...)

	// Append the stack trace when wrapping the errors
	if Log.Errors.StackTrace {
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

	// pad space between elements
	PrintData = tools.PadSpaceArray(PrintData)

	// add newline at the end
	if NewlineConfig.Auto {
		// check if the last element is a string and ends with a newline
		if NewlineConfig.Smart && len(c.Data) > 0 {
			if StrValue, ok := PrintData[len(PrintData)-1].(string); ok {
				if len(StrValue) > 0 && StrValue[len(StrValue)-1] != '\n' {
					PrintData = append(PrintData, "\n")
				}
			} else {
				// if the last element is not a string, just append a newline
				PrintData = append(PrintData, "\n")
			}
		} else {
			// if smart newline is disabled, just append a newline
			PrintData = append(PrintData, "\n")
		}
	}

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
	if Log.RemoveColor {
		PrintData = RemoveColorData
	}

	return PrintData, RemoveColorData
}
