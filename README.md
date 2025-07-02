# Noa
A Golang logging library that supports basic log printing, recording, automatic log rotation, and more.  
It can be quickly integrated into existing projects and offers flexible configuration options and extensibility.

## Installation
```bash
go get -u github.com/noa-log/noa
```

## Quick Start
```go
package main

import (
    "errors"
    "github.com/noa-log/noa"
)

func main() {
    // Create a new logger instance
    logger := noa.NewLog()
    // Configure the logger, such as setting the log level, format, etc.
    logger.Level = noa.INFO

    // Print logs
    logger.Debug("Test", "This is a Debug log")
    logger.Info("Test", "This is an Info log")
    logger.Warning("Test", "This is a Warn log")
    logger.Error("Test", "This is an Error log")
    logger.Fatal("Test", "This is a Fatal log")

    // Print an error
    err := errors.New("An example error")
    logger.Error("Test", err)
}
```
For more usage details and configuration options, please refer to the [Documentation](docs/en/README.md)

## Features
- Supports multiple log levels
- Supports automatic log rotation based on time
- Supports log formatting
- Supports handle before and after printing logs, allowing for modification or reporting of log information
- Provides extensive configuration options to customize logging behavior
- Supports integration with some third-party libraries such as `Gin`, `Gorm`, etc.
- Enhances error context information by wrapping errors

## Integrations
Here are some officially maintained integration libraries that provide Noa support for popular frameworks and libraries:
- [noa-gin](https://github.com/noa-log/noa-gin) - Integrate Noa with the Gin framework
- [noa-echo](https://github.com/noa-log/noa-echo) - Integrate Noa with the Echo framework
- [noa-gorm](https://github.com/noa-log/noa-gorm) - Integrate Noa with Gorm
- [noa-sentry](https://github.com/noa-log/noa-sentry) - Integrate Noa with Sentry

## License
This project is open-sourced under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0). Please comply with the terms when using it.