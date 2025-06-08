# Noa
一个Golang的日志库，支持基础日志打印，记录，自动分割日志等功能。    
可以快速集成到现有项目中，提供了灵活的配置选项和扩展性。

## 安装
```bash
go get -u github.com/noa-log/noa
```

## 快速开始
```go
package main

import (
    "errors"
    "github.com/noa-log/noa"
)

func main() {
    // 创建一个新的日志实例
    logger := noa.NewLog()
    // 配置日志实例，如设置日志级别、格式等
    logger.Level = noa.INFO

    // 打印日志
    logger.Debug("Test", "这是一条Debug日志")
    logger.Info("Test", "这是一条Info日志")
    logger.Warning("Test", "这是一条Warn日志")
    logger.Error("Test", "这是一条Error日志")
    logger.Fatal("Test", "这是一条Fatal日志")

    // 打印错误
    err := errors.New("一个错误示例")
    logger.Error("Test", err)
}
```
更多使用方法和配置选项，请参考[文档](docs/cn/README.md)

## 功能
- 支持多种日志级别
- 支持按时间自动分割日志
- 支持日志格式化
- 支持打印前后插入中间件，方便修改或上报日志信息
- 提供大量的配置选项用于定制日志行为
- 支持集成到部分第三方库中，如`Gin`、`Gorm`等
- 通过将错误进行包装，提供更好的错误上下文信息

## 许可
本项目基于[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)协议开源。使用时请遵守协议的条款。