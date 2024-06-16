package log

import "fmt"

// log pkg is interface foe EQLDB logger
type log interface {
	Log(message string, verbosity LogLevel)
}

// LogLevel represent the valid logging verbosity levels
type LogLevel uint8

const (
	//LogOff is for logging nothing
	LogOff LogLevel = iota
	//LoInfo is for Logging Informatice events. This the default logging level
	LogInfo
	// LogDebug is for logging information useful for closely tracing the operation of the EQLDB.
	LogDebug
)

type eqldbLogger struct {
	logger    log
	verbosity LogLevel
}

type defaultLogger struct{}

// Log the message uisng the built-in Golang logging package
func (logger defaultLogger) Log(message string, verbosity LogLevel) {
	fmt.Printf("log: %v\n", message)
}

func (eqldbLogger *eqldbLogger) log(verbosityLevel LogLevel, message string) {
	if verbosityLevel <= eqldbLogger.verbosity {
		switch verbosityLevel {
		case LogInfo:
			eqldbLogger.logger.Log("[INFO]"+message, verbosityLevel)
		case LogDebug:
			eqldbLogger.logger.Log("[DEBUG]"+message, verbosityLevel)
		default:
			eqldbLogger.logger.Log(message, verbosityLevel)
		}
	}
}

func (eqldbLogger *eqldbLogger) logf(verbosityLevel LogLevel, message string, args ...interface{}) {
	if verbosityLevel <= eqldbLogger.verbosity {
		switch verbosityLevel {
		case LogInfo:
			eqldbLogger.logger.Log(fmt.Sprintf("[INFO] "+message, args...), verbosityLevel)
		case LogDebug:
			eqldbLogger.logger.Log(fmt.Sprintf("[DEBUG] "+message, args...), verbosityLevel)
		default:
			eqldbLogger.logger.Log(fmt.Sprintf(message, args...), verbosityLevel)
		}
	}
}
