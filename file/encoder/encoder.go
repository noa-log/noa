/*
 * @Author: nijineko
 * @Date: 2025-06-10 11:58:44
 * @LastEditTime: 2025-06-10 12:38:19
 * @LastEditors: nijineko
 * @Description: noa file writer encoder package
 * @FilePath: \noa\file\encoder\encoder.go
 */
package encoder

import (
	"os"
)

type Encoder interface {
	/**
	* @description: Write data to file
	* @param {*os.File} FileHandle file handle
	* @param {[]any} PrintData data to encode
	* @return {error} error
	 */
	Write(FileHandle *os.File, PrintData []any) error
}
