/*
 * @Author: nijineko
 * @Date: 2025-06-08 14:38:51
 * @LastEditTime: 2025-06-08 14:52:59
 * @LastEditors: nijineko
 * @Description: log handle test package
 * @FilePath: \noa\handle_test.go
 */
package noa

import (
	"fmt"
	"os"
	"testing"
)

func TestBeforeHandle(t *testing.T) {
	Log := NewLog()

	Log.AddBeforeHandle(func(Level *int, Source *string, Data ...*any) error {
		*Level = WARNING
		*Source = "Hijack"
		for Index := range Data {
			if Index == 0 && Data[Index] != nil {
				*Data[Index] = "This log has been hijacked"
			}
		}

		return nil
	})

	Log.Info("Test", "This is an info log")
}

func TestAfterHandle(t *testing.T) {
	Log := NewLog()

	Log.AddAfterHandle(func(Level int, Source string, Data ...any) error {
		if Level == ERROR {
			os.WriteFile("errors.log", []byte(fmt.Sprintf("%s: %v\n", Source, Data[0])), 0644)
		}
		return nil
	})

	Log.Error("Test", "This is an error log")
}
