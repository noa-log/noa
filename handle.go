/*
 * @Author: nijineko
 * @Date: 2025-06-08 10:52:00
 * @LastEditTime: 2025-06-08 13:14:48
 * @LastEditors: nijineko
 * @Description: log handle package
 * @FilePath: \noa\handle.go
 */
package noa

type BeforeHandleFunc func(Level *int, Source *string, Data ...*any) error
type AfterHandleFunc func(Level int, Source string, Data ...any) error

/**
 * @description: Add a handle function to run before logging
 * @description: if error returned, will panic
 * @param {BeforeHandleFunc} HandleFun
 */
func (l *LogConfig) AddBeforeHandle(HandleFunc BeforeHandleFunc) {
	l.beforeHandles = append(l.beforeHandles, HandleFunc)
}

/**
 * @description: Add a handle function to run after logging
 * @description: if error returned, will panic
 * @param {AfterHandleFunc} HandleFunc
 */
func (l *LogConfig) AddAfterHandle(HandleFunc AfterHandleFunc) {
	l.bfterHandles = append(l.bfterHandles, HandleFunc)
}
