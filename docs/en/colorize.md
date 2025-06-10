# Add Color and Formatting to Text
Noa uses the `github.com/noa-log/colorize` library to add color and formatting to text. You can control whether to remove color and formatting before printing by configuring the `RemoveColor` property of the log instance.

## Usage Example
```go
package main

import (
    "github.com/noa-log/noa"
    "github.com/noa-log/colorize"
)

func main() {
    // Create a new log instance
    logger := noa.NewLog()

    // Print text with color
    logger.Info("Test", colorize.RedText("This is a red text"))
}
```

For more information about the `github.com/noa-log/colorize` package, please refer to its [Documentation](https://github.com/noa-log/colorize/blob/main/README.md)