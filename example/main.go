package main

import (
	"log"

	"github.com/speedtrade/go-utils/config"
	"github.com/speedtrade/go-utils/flags"
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
	config.Init(flags.BaseConfigPath() + "/" + flags.Env())

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
				ConsoleLoggingEnabled: conf.GetBool(LoggerConsoleEnabled),
				Filename:              conf.GetString(LoggerFileName),
				MaxSize:               conf.GetInt(LoggerMaxSizeInMB),
				MaxAge:                conf.GetInt(LoggerMaxAgeInDays),
				MaxBackups:            conf.GetInt(LoggerMaxBackUps),
				Compress:              conf.GetBool(LoggerCompress),
			},
			Lef: func(lvl logger.Level) bool {
				return lvl <= logger.InfoLevel
			},
		},
		{
			Logopt: model.LoggerOptions{
				ConsoleLoggingEnabled: conf.GetBool(LoggerConsoleEnabled),
				Filename:              conf.GetString(LoggerErrorFileName),
				MaxSize:               conf.GetInt(LoggerMaxSizeInMB),
				MaxAge:                conf.GetInt(LoggerMaxAgeInDays),
				MaxBackups:            conf.GetInt(LoggerMaxBackUps),
				Compress:              conf.GetBool(LoggerCompress),
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
