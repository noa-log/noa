/*
 * @Author: nijineko
 * @Date: 2025-06-10 11:59:30
 * @LastEditTime: 2025-06-10 12:38:28
 * @LastEditors: nijineko
 * @Description: text log encoder
 * @FilePath: \noa\file\encoder\text\text.go
 */
package text

import (
	"fmt"
	"os"

	"github.com/noa-log/noa/tools"
)

// TextEncoder struct
type TextEncoder struct{}

/**
 * @description: Create a new TextEncoder instance
 * @return {*TextEncoder} TextEncoder instance
 */
func NewTextEncoder() *TextEncoder {
	return &TextEncoder{}
}

/**
 * @description: Text encode method
 * @param {*os.File} FileHandle file handle
 * @param {[]any} PrintData data to encode
 * @return {[]byte} encoded data
 * @return {error} error
 */
func (te *TextEncoder) Write(FileHandle *os.File, PrintData []any) error {
	if _, err := fmt.Fprint(FileHandle, tools.PadSpaceArray(PrintData)...); err != nil {
		return err
	}

	return nil
}
