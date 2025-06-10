/*
 * @Author: nijineko
 * @Date: 2025-06-10 12:12:04
 * @LastEditTime: 2025-06-10 12:12:05
 * @LastEditors: nijineko
 * @Description: pad space util
 * @FilePath: \noa\tools\padSpace.go
 */
package tools

/**
 * @description: Pad space between elements in an array.
 * @param {[]any} Data Array of data to pad with spaces.
 * @return {[]any} New array with spaces padded between elements.
 */
func PadSpaceArray(Data []any) []any {
	ResultData := make([]any, 0, len(Data)*2-1)
	for Index, Value := range Data {
		ResultData = append(ResultData, Value)
		if Index < len(Data)-1 {
			// If the value is a string and ends with a newline, skip adding a space
			if StrValue, ok := Value.(string); ok && (len(StrValue) > 0 && StrValue[len(StrValue)-1] == '\n') {
				continue
			}

			ResultData = append(ResultData, " ")
		}
	}
	return ResultData
}
