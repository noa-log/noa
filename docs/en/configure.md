# Configure Log Instance
Noa supports highly flexible configuration options, allowing you to set the behavior of logging instances via code or configuration files.

## Configuration Options
|      Option       |       Description        |    Default Value    |                                        Notes                                         |
| :---------------: | :----------------------: | :-----------------: | :----------------------------------------------------------------------------------: |
|       Level       |        Log level         |      noa.DEBUG      | Supports `noa.DEBUG`, `noa.INFO`, `noa.WARNING`, `noa.ERROR`, `noa.FATAL`, `noa.OFF` |
|    RemoveColor    |       Remove color       |        false        |                     Removes all color parameters before printing                     |
|    TimeFormat     |     Timestamp format     | 2006-01-02 15:04:05 |                                Uses Go's time format                                 |
| Errors.StackTrace | Print error stack trace  |        true         |                                                                                      |
| Errors.CallerSkip |   Stack depth to skip    |          3          |                                                                                      |
|   Writer.Enable   | Enable log file writing  |        true         |                                                                                      |
| Writer.FolderPath |   Log file folder path   |       ./logs        |                                                                                      |
| Writer.TimeFormat | Log filename time format |     2006-01-02      |                Uses Go's time format, affects automatic file rotation                |
|   Encoder.Print   |    Log output encoder    |  NewTextEncoder()   |            Sets the encoder for console output; defaults to text encoder             |
|   Encoder.Write   |   Log writing encoder    |  NewTextEncoder()   |              Sets the encoder for file output; defaults to text encoder              |

## Example of Modifying Configuration
```go
package main

import (
    "github.com/noa-log/noa"
)

func main() {
    // Create a new logger instance
    logger := noa.NewLog()

    // Set log level to INFO
    logger.Level = noa.INFO
    // Remove log colors
    logger.RemoveColor = true
    // Set timestamp format for log prefix
    logger.TimeFormat = "2006-01-02 15:04:05"
    // Enable log file writing
    logger.Writer.Enable = true
    // ...
}
```