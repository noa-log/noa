# Configure Scheduled Tasks
Noa uses the `github.com/noa-log/noa-timer` package to provide scheduled task functionality. You can import and use it as follows:

## Installation
```bash
go get -u github.com/noa-log/noa-timer
```

## Usage Example
```go
package main

import (
    "github.com/noa-log/noa"
    noatimer "github.com/noa-log/noa-timer"
)

func main() {
    // Create a new log instance
    logger := noa.NewLog()

    // Register default timer tasks: compress logs at 00:30 and clear logs older than 7 days at 00:35 every day
    go noatimer.StartDefaultTask(logger)

    logger.Info("Test", "Starting Noa Timer")

    // ... Execute other business logic
}
```

For more information about the `github.com/noa-log/noa-timer` package, please refer to its [Documentation](https://github.com/noa-log/noa-timer/blob/main/README.md)