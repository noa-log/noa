# Print Logs
Noa supports quickly printing logs at various levels. You can also use the `Print()` methods to print custom logs.

## Quick Log Print
```go
package main

import (
    "github.com/noa-log/noa"
)

func main() {
    // Create a new log instance
    logger := noa.NewLog()

    // Print logs at different levels
    logger.Debug("Test", "This is a Debug log")
    logger.Info("Test", "This is an Info log")
    logger.Warning("Test", "This is a Warning log")
    logger.Error("Test", "This is an Error log")
    logger.Fatal("Test", "This is a Fatal log")
}
```

## Printing Errors
Noa supports printing error messages. When using `Error()` or `Fatal()` methods, it can automatically extract the error stack trace and wrap it with the `errors.Error` type. The stack trace will be printed automatically.
```go
package main

import (
    "errors"
    "github.com/noa-log/noa"
)

func main() {
    // Create a new log instance
    logger := noa.NewLog()

    // Print an error
    err := errors.New("An example error")
    logger.Error("Test", err)
}
```

## Printing Custom Logs
Noa supports using the `Print()` methods to print custom log messages.
```go
package main

import (
    "errors"
    "github.com/noa-log/noa"
)

func main() {
    // Create a new log instance
    logger := noa.NewLog()

    // Print custom logs
    logger.Print(noa.INFO, "Test", "This is a custom log")
    logger.Print(noa.INFO, "Test", "This is a custom log with a newline")

    // Bypass error wrapping and print the error directly
    err := errors.New("An example error")
    logger.Print(noa.ERROR, "Test", err)
}
```