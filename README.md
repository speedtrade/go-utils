# go-utils

Go utils contains the set of common reusable utility methods which can be reused across all Go projects.

## Installation

To install `go-utils` package, you need to install Go and set your Go workspace first.

1. The first need Go installed (version 1.17+ is required), then you can use the below Go command to install go-utils.

```shell
go get github.com/speedtrade/go-utils
```
4. Import it in your code:

```go
package main

import "github.com/speedtrade/go-utils"
```

## Usage

### Reading Config

Config file is located at resources/dev directory and name is config.yml.
```yml
logger:
  consoleLoggingEnabled: true
  fileName: logs/access.log
  errorFileName: logs/error.log
  maxSizeInMB: 1024
  maxBackups: 10
  maxAgeInDays: 2
  compress: true
```
Go code to read above config.yml file
```go
package main

import (
	"log"

	"github.com/speedtrade/go-utils/config"
)

const (
	LoggerConsoleEnabled = "logger.consoleLoggingEnabled"
	LoggerFileName       = "logger.fileName"
	LoggerErrorFileName  = "logger.errorFileName"
	LoggerMaxSizeInMB    = "logger.maxSizeInMB"
	LoggerMaxBackUps     = "logger.maxBackups"
	LoggerMaxAgeInDays   = "logger.maxAgeInDays"
	LoggerCompress       = "logger.compress"
)

func main() {
	config.Init("resources/dev")

	conf, err := config.Get("config")

	if err != nil {
		log.Fatal("error getting config err : ", err)
	}

	log.Println("Logger Filename : ", conf.GetString(LoggerFileName))
	log.Println("Logger Error Filename : ", conf.GetString(LoggerErrorFileName))
	log.Println("Logger Max Size In MB : ", conf.GetString(LoggerMaxSizeInMB))
	log.Println("Logger Max Back Ups : ", conf.GetString(LoggerMaxBackUps))
}
```

### Logger

The logger is based on uber-zap module.

```go
package main

import (
	"log"

	"github.com/speedtrade/go-utils/config"
	"github.com/speedtrade/go-utils/logger"
	"github.com/speedtrade/go-utils/model"
)

const (
	LoggerConsoleEnabled = "logger.consoleLoggingEnabled"
	LoggerFileName       = "logger.fileName"
	LoggerErrorFileName  = "logger.errorFileName"
	LoggerMaxSizeInMB    = "logger.maxSizeInMB"
	LoggerMaxBackUps     = "logger.maxBackups"
	LoggerMaxAgeInDays   = "logger.maxAgeInDays"
	LoggerCompress       = "logger.compress"
)

func main() {
	config.Init("resources/dev")

	conf, err := config.Get("config")

	if err != nil {
		log.Fatal("error getting config err : ", err)
	}

	log.Println("Logger Filename : ", conf.GetString(LoggerFileName))
	log.Println("Logger Error Filename : ", conf.GetString(LoggerErrorFileName))
	log.Println("Logger Max Size In MB : ", conf.GetString(LoggerMaxSizeInMB))
	log.Println("Logger Max Back Ups : ", conf.GetString(LoggerMaxBackUps))

	var tops = []logger.TeeOption{
		{

			Logopt: model.LoggerOptions{
				Filename:   conf.GetString(LoggerFileName),
				MaxSize:    conf.GetInt(LoggerMaxSizeInMB),
				MaxAge:     conf.GetInt(LoggerMaxAgeInDays),
				MaxBackups: conf.GetInt(LoggerMaxBackUps),
				Compress:   conf.GetBool(LoggerCompress),
			},
			Lef: func(lvl logger.Level) bool {
				return lvl <= logger.InfoLevel
			},
		},
		{
			Logopt: model.LoggerOptions{
				Filename:   conf.GetString(LoggerErrorFileName),
				MaxSize:    conf.GetInt(LoggerMaxSizeInMB),
				MaxAge:     conf.GetInt(LoggerMaxAgeInDays),
				MaxBackups: conf.GetInt(LoggerMaxBackUps),
				Compress:   conf.GetBool(LoggerCompress),
			},
			Lef: func(lvl logger.Level) bool {
				return lvl > logger.InfoLevel
			},
		},
	}

	log := logger.NewTeeWithRotate(tops)
	logger.ResetDefault(log)

	for i := 0; i < 20; i++ {
		logger.Info("example:", logger.String("app", "start ok"), logger.Int("major version", 3))
		logger.Error("example:", logger.String("app", "crash"), logger.Int("reason", -1))
	}
}
```

### Logger Ref.
https://www.sobyte.net/post/2022-03/uber-zap-advanced-usage
