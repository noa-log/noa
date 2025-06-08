# Wrap Errors
Noa provides the `errors` package to wrap errors, allowing for better contextual information. Wrapped errors can automatically include stack trace information when printed.

## Manually Wrapping Errors
```go
package main

import (
    baseErrors "errors"
    "github.com/noa-log/noa/errors"
)

func main() {
    // Create a new logger instance
    logger := noa.NewLog()

    // Manually wrap an error
    err := errors.Wrap(baseErrors.New("an example error"), 3)
    
    // Print the wrapped error
    logger.Println(noa.ERROR, "Test", err)
}
```