# Set Encoder
Noa supports custom encoders to control the output format of logs. You can set an encoder using the `SetEncoder()` method or configure different encoders for printing and writing logs using the `Encoder.Print` and `Encoder.Write` fields.  
By default, Noa uses a built-in text encoder for both printing and writing logs.

## Using JSON Encoder
Noa officially provides a JSON encoder implementation, which you can import and use as follows:

### Installation
```bash
go get -u github.com/noa-log/noa-encoder-json
```

### Usage
```go
package main

import (
    "errors"
    "github.com/noa-log/noa"
    noaencoderjson "github.com/noa-log/noa-encoder-json"
)

func main() {
    // Create a new log instance
    logger := noa.NewLog()
    // Set the encoder to the JSON encoder
    logger.SetEncoder(noaencoderjson.NewJSONEncoder(logger))

    // Print Log
    logger.Info("Test", "This is an info message")
    logger.Error("Test", errors.New("an example error"))
}
```

## Setting Different Encoders
You can also configure different encoders for printing and writing. For example, you can use a text encoder for printing and a JSON encoder for writing:
```go
package main

import (
    "errors"
    "github.com/noa-log/noa"
    noaencoderjson "github.com/noa-log/noa-encoder-json"
)

func main() {
    // Create a new log instance
    logger := noa.NewLog()
    
    // Set the print encoder to the text encoder
    logger.Encoder.Print = noa.NewTextEncoder(logger)
    // Set the write encoder to the JSON encoder
    logger.Encoder.Write = noaencoderjson.NewJSONEncoder(logger)

    // Log messages
    logger.Info("Test", "This is an info message")
    logger.Error("Test", errors.New("an example error"))
}
```

## Custom Encoder
If you want to implement your own encoder, you can do so by implementing the `encoder.Encoder` interface from the `github.com/noa-log/noa/encoder` package. Here's a basic example of a custom encoder:

**Note:** When implementing a custom encoder, try to support Noa’s configuration options such as RemoveColor, TimeFormat, Errors.StackTrace, etc., to ensure expected behavior.
```go
package main

import (
	"fmt"
	"os"

	"github.com/noa-log/colorize"
	"github.com/noa-log/noa"
	"github.com/noa-log/noa/encoder"
	"github.com/noa-log/noa/errors"
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

// Implement the Print method
func (mce *MyCustomEncoder) Print(c *encoder.Context) {
	printData := []any{
		c.Time.Format(mce.logger.TimeFormat),
		c.Level,
		"<" + c.Source + ">",
	}

	data := c.Data
	// Add error stack trace information
	if mce.logger.Errors.StackTrace {
		errorStackData := make([]any, 0, len(data))
		for _, value := range data {
			if wrapError, ok := value.(*errors.Error); ok {
				errorStackData = append(errorStackData, value, "\n", wrapError.StackFormat())
			} else {
				errorStackData = append(errorStackData, value)
			}
		}
		data = errorStackData
	}

	// Remove color codes
	if mce.logger.RemoveColor {
		for index, value := range data {
			if valueStr, ok := value.(string); ok {
				data[index] = colorize.Remove(valueStr)
			}
		}
	}

	printData = append(printData, data...)

	// Cache print data
	c.Set("printData", printData)

	fmt.Println(printData...)
}

// Implement method to get the file extension for written logs
func (mce *MyCustomEncoder) WriteFileExtension() string {
	return ".log"
}

// Implement the Write method
func (mce *MyCustomEncoder) Write(FileHandle *os.File, c *encoder.Context) error {
	// Attempt to retrieve cached print data from the context
	printData := c.Get("printData")
	printDataSlice, ok := printData.([]any)
	if printData == nil || !ok {
		return fmt.Errorf("printData is nil or not a slice: %v", printData)
	}

	if _, err := fmt.Fprintln(FileHandle, printDataSlice...); err != nil {
		return err
	}

	return nil
}

func main() {
	// Create a new log instance
	logger := noa.NewLog()
	// Set custom encoder
	logger.SetEncoder(NewMyCustomEncoder(logger))

	// Log a message
	logger.Info("Test", "This is an info message")
}
```

### Encoder Interface Signature
```go
type Encoder interface {
	/**
	* @description: print log data
	* @param {*Context} c encoder context
	 */
	Print(c *Context)
	/**
	* @description: Get file extension for the encoded file (Note: A file can only have one extension, such as .log. Having multiple extensions may cause parsing errors.)
	* @return {string} file extension
	 */
	WriteFileExtension() string
	/**
	* @description: Write data to file
	* @param {*os.File} FileHandle file handle
	* @param {*Context} c encoder context
	* @return {error} error
	 */
	Write(FileHandle *os.File, c *Context) error
}
```