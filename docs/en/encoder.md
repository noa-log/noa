# Set Write Encoder
Noa supports setting a custom write encoder, which can be configured via the `Writer.Encoder` option of the log instance. By default, Noa uses its built-in Text encoder.

## Switching to JSON Encoder
Noa officially provides a JSON encoder implementation that you can import and use as follows:

### Installation
```bash
go get -u github.com/noa-log/noa-encoder-json
```

```go
package main

import (
    "errors"
    "github.com/noa-log/noa"
    noaencoderjson "github.com/noa-log/noa-encoder-json"
)

func main() {
    // Create a new logger instance
    logger := noa.NewLog()
    // Set the encoder to the JSON encoder
    logger.Writer.Encoder = noaencoderjson.NewJSONEncoder(logger)

    // Print Log
    logger.Info("Test", "This is an info message")
    logger.Error("Test", errors.New("an example error"))
}
```

## Writing a Custom Encoder
If you want to write your own encoder, you can do so by implementing the `encoder.Encoder` interface from the `github.com/noa-log/noa/file/encoder` package. Here is a simple example of a custom encoder:
```go
package main

import (
	"fmt"
	"os"
	"github.com/noa-log/noa"
	"github.com/noa-log/noa/tools"
    "github.com/noa-log/noa/tools/output"
)

type MyCustomEncoder struct {
	logger *noa.LogConfig
}

// Create a new custom encoder
func NewMyCustomEncoder(logger *noa.LogConfig) *MyCustomEncoder {
	return &MyCustomEncoder{
		logger: logger,
	}
}

// Implement the method to get the file extension
func (mce *MyCustomEncoder) FileExtension() string {
	return ".log"
}

// Implement the write method
func (mce *MyCustomEncoder) Write(FileHandle *os.File, PrintData []any) error {
    // Use the `output.UnwrapPrintData` utility method to parse `PrintData` and extract time, level, source, etc.
	time, level, source, printData := output.UnwrapPrintData(mce.logger, PrintData)

	if _, err := fmt.Fprint(FileHandle, tools.PadSpaceArray([]any{time, level, source, printData})...); err != nil {
		return err
	}

	return nil
}

func main() {
	// Create a new log instance
	logger := noa.NewLog()
	// Set the custom encoder
	logger.Writer.Encoder = NewMyCustomEncoder(logger)

	// Print Log
	logger.Info("Test", "This is an info message")
}
```