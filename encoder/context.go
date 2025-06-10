/*
 * @Author: nijineko
 * @Date: 2025-06-10 21:56:37
 * @LastEditTime: 2025-06-10 21:56:43
 * @LastEditors: nijineko
 * @Description: noa encoder context package
 * @FilePath: \noa\encoder\context.go
 */
package encoder

// encoder context struct
type Context struct {
	Level  int    // Log level
	Source string // Log source
	Data   []any  // Datas

	ExtendKeys map[string]any // extend keys for custom data
}

/**
 * @description: Create a new encoder context
 * @param {int} Level log level
 * @param {string} Source log source
 * @param {[]any} Data data to log
 * @return {*Context} new encoder context
 */
func NewContext(Level int, Source string, Data []any) *Context {
	return &Context{
		Level:      Level,
		Source:     Source,
		Data:       Data,
		ExtendKeys: make(map[string]any),
	}
}

/**
 * @description: Set custom data in the encoder context
 * @param {string} Key
 * @param {any} Value
 */
func (c *Context) Set(Key string, Value any) {
	c.ExtendKeys[Key] = Value
}

/**
 * @description: Get custom data from the encoder context
 * @param {string} Key
 * @return {any} Value
 */
func (c *Context) Get(Key string) any {
	if Value, ok := c.ExtendKeys[Key]; ok {
		return Value
	}
	return nil
}
