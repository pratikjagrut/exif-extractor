package log

import (
	"fmt"
	"log"
	"os"
)

var (
	EnableLogs = true // Flag to enable or disable logs
)

const (
	// Define color codes
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

func Info(format string, v ...interface{}) {
	if EnableLogs {
		log.Printf(InfoColor, fmt.Sprintf(format, v...))
	}
}

func Warning(format string, v ...interface{}) {
	if EnableLogs {
		log.Printf(WarningColor, fmt.Sprintf(format, v...))
	}
}

func Error(format string, v ...interface{}) {
	if EnableLogs {
		log.Printf(ErrorColor, fmt.Sprintf(format, v...))
	}
}

func Fatal(format string, v ...interface{}) {
	Error(format, v...)
	os.Exit(1)
}
