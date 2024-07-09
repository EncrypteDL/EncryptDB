package driver

import (
	"fmt"
	"log"
)

type Logger interface {
	//Log the message uisng the build-in Golang logging pakg
	Log(message string, verbosity LogLevel)
}

type LogLevel uint8

// //New Contsruct a new log for uses
// func New(w io.Writer, logLevel LogLevel, servviceName string, verboity LogLevel) *Logger{
// 	return New(eqldbLogger)
// }

const (
	// LogOff is for logging nothing.
	LogOff LogLevel = iota
	// LogInfo is for logging informative events. This is the default logging level.
	LogInfo
	// LogDebug is for logging information useful for closely tracing the operation of the QLDBDriver.
	LogDebug
)

type eqldbLogger struct {
	logger    Logger
	verbosity LogLevel
}

func (qldbLogger *eqldbLogger) log(verbosityLevel LogLevel, message string) {
	if verbosityLevel <= qldbLogger.verbosity {
		switch verbosityLevel {
		case LogInfo:
			qldbLogger.logger.Log("[INFO] "+message, verbosityLevel)
		case LogDebug:
			qldbLogger.logger.Log("[DEBUG] "+message, verbosityLevel)
		default:
			qldbLogger.logger.Log(message, verbosityLevel)
		}
	}
}

func (qldbLogger *eqldbLogger) ogf(verbosityLevel LogLevel, message string, args ...interface{}) {
	if verbosityLevel <= qldbLogger.verbosity {
		switch verbosityLevel {
		case LogInfo:
			qldbLogger.logger.Log(fmt.Sprintf("[INFO] "+message, args...), verbosityLevel)
		case LogDebug:
			qldbLogger.logger.Log(fmt.Sprintf("[DEBUG] "+message, args...), verbosityLevel)
		default:
			qldbLogger.logger.Log(fmt.Sprintf(message, args...), verbosityLevel)
		}
	}
}

type defaultLogger struct{}

// Log the message using the built-in Golang logging package.
func (logger defaultLogger) Log(message string, verbosity LogLevel) {
	log.Println(message)
}
