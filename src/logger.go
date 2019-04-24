package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// fileLog is to log information to a file
var fileLog = log.New()

// consoleLog is to log information to console.
var consoleLog = log.New()

// LogAll is to log in both console and in log file
func LogAll() {

	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

}
