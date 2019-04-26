package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// To log information to a file
var logAPI = log.New()

// LoadLogger is to set logger
func LoadLogger() {
	// Check if API is in production
	// If true, log to file
	// else, log to console
	if viper.GetBool("production") {
		logAPI.SetLevel(log.InfoLevel)
		// Log as JSON instead of the default ASCII formatter.
		logAPI.SetFormatter(&log.JSONFormatter{})
		// Log to a file
		// Format MM-DD-YYYY
		file, err := os.OpenFile(viper.GetString("logAddr")+time.Now().Format("01-02-2006")+".log", os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			logAPI.SetOutput(file)
		} else {
			logAPI.Info("Failed to log to file, using default stderr")
		}
	} else {
		logAPI.SetLevel(log.TraceLevel)
		// Log as default ASCII formatter.
		logAPI.SetFormatter(&log.TextFormatter{})
		// Output to stdout instead of the default stderr
		logAPI.SetOutput(os.Stdout)
	}
}

// Log is to log according to the environment
func Log(tag string, message string, fields ...log.Fields) {
	field := log.Fields{}
	if len(fields) > 0 {
		field = fields[0]
	}
	context := logAPI.WithFields(field)
	switch tag {
	// Something very low level.
	case "trace":
		context.Trace(message)
	// Useful debugging information.
	case "debug":
		context.Debug(message)
	// Something noteworthy happened!")
	case "info":
		context.Info(message)
	// You should probably take a look at this.
	case "warn":
		context.Warn(message)
	// Something failed but I'm not quitting.
	case "error":
		context.Error(message)
	// Calls os.Exit(1) after logging
	case "fatal":
		context.Fatal(message)
		// Calls panic() after logging
	case "panic":
		context.Panic(message)
	default:
		context.Debug(message)
	}
}
