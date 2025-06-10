/*
 * @Author: nijineko
 * @Date: 2025-06-10 11:58:44
 * @LastEditTime: 2025-06-10 13:28:34
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
	* @description: Get file extension for the encoded file (Note: A file can only have one extension, such as .log. Having multiple extensions may cause parsing errors.)
	* @return {string} file extension
	 */
	FileExtension() string
	/**
	* @description: Write data to file
	* @param {*os.File} FileHandle file handle
	* @param {[]any} PrintData data to encode
	* @return {error} error
	 */
	Write(FileHandle *os.File, PrintData []any) error
}
