# Registering Handles
Noa supports the registration of handles, allowing you to insert custom logic before and after logging. This makes it easier to modify or report log information.

**Note**: Handle functions are executed in the order they are registered. Only the before handles can modify the log data.

## Example of Registering Handles
```go
package main

import (
    "fmt"
    "os"

    "github.com/noa-log/noa"
)

func main() {
    // Create a new logger instance
    logger := noa.NewLog()

    // Register a before-handle, used to modify log data
    logger.AddBeforeHandle(func(Level *int, Source *string, Data ...*any) error {
        *Level = noa.WARNING
        *Source = "Hijack"
        for Index := range Data {
            if Index == 0 && Data[Index] != nil {
                *Data[Index] = "This log has been hijacked"
            }
        }

        return nil
    })

    // Register an after-handle, used to report log data
    logger.AddAfterHandle(func(Level int, Source string, Data ...any) error {
        if Level == noa.ERROR {
            os.WriteFile("errors.log", []byte(fmt.Sprintf("%s: %v\n", Source, Data[0])), 0644)
        }
        return nil
    })

    // Print a log
    logger.Info("Test", "This is an Info log")
}
```

## Handle Function Signatures
```go
type BeforeHandleFunc func(Level *int, Source *string, Data ...*any) error
type AfterHandleFunc func(Level int, Source string, Data ...any) error
```