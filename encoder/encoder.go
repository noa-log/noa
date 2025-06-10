/*
 * @Author: nijineko
 * @Date: 2025-06-10 21:23:29
 * @LastEditTime: 2025-06-10 21:33:31
 * @LastEditors: nijineko
 * @Description: noa encoder package
 * @FilePath: \noa\encoder\encoder.go
 */
package encoder

import "os"

// Encoder interface
type Encoder interface {
	/**
	* @description: print log data
	* @param {*Context} c encoder context
	 */
	Print(c *Context)
	/**
	* @description: Get file extension for the encoded file (Note: A file can only have one extension, such as .log. Having multiple extensions may cause parsing errors.)
	* @return {string} file extension
	 */
	WriteFileExtension() string
	/**
	* @description: Write data to file
	* @param {*os.File} FileHandle file handle
	* @param {*Context} c encoder context
	* @return {error} error
	 */
	Write(FileHandle *os.File, c *Context) error
}
