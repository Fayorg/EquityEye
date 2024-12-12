package logs

import (
	"fmt"
	"log"
	"os"
)

// LogLevel defines the severity of the log message
type LogLevel int

const (
	INFO LogLevel = iota
	MESSAGE
	WARN
	ERROR
)

// Color codes
const (
	reset  = "\033[0m"
	blue   = "\033[34m"
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
)

// Package-level logger instance
var logger *log.Logger

func init() {
	// Initialize the logger during package initialization
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

// Log logs the message with the appropriate level and color
func Log(level LogLevel, format string, v ...interface{}) {
	var color string
	var prefix string

	switch level {
	case INFO:
		color = blue
		prefix = "[INFO] "
	case MESSAGE:
		color = green
		prefix = "[MESSAGE] "
	case WARN:
		color = yellow
		prefix = "[WARN] "
	case ERROR:
		color = red
		prefix = "[ERROR] "
	default:
		color = reset
		prefix = "[LOG] "
	}

	logger.Printf("%s%s%s%s", color, prefix, fmt.Sprintf(format, v...), reset)
}

// Convenience methods for specific levels
func Info(format string, v ...interface{}) {
	Log(INFO, format, v...)
}

func Message(format string, v ...interface{}) {
	Log(MESSAGE, format, v...)
}

func Warn(format string, v ...interface{}) {
	Log(WARN, format, v...)
}

func Error(format string, v ...interface{}) {
	Log(ERROR, format, v...)
}
