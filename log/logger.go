package log

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// Error logs a message as of error type.
func Error(message string) {
	log("ERROR", message)
}

// Debug logs a message as of debug type.
func Debug(message string) {
	log("DEBUG", message)
}

// Info logs a message as of information type.
func Info(message string) {
	log("INFO", message)
}

// Warn logs a message as of warning type.
func Warn(message string) {
	log("WARN", message)
}

// Theme logs a message as of theme type.
func Theme(message string) {
	log("THEME", message)
}

// Default logs a message as of default type.
func Default(message string) {
	log("", message)
}

// log logs a message with colour.
func log(level string, message string) {

	var av aurora.Value

	switch level {
	case "ERROR":
		av = aurora.Red(message)
		break
	case "DEBUG":
		av = aurora.Green(message)
		break
	case "INFO":
		av = aurora.Yellow(message)
		break
	case "WARN":
		av = aurora.Brown(message)
		break
	case "THEME":
		av = aurora.Blue(message)
		break
	default:
		fmt.Println(message)
		break
	}

	fmt.Println(av)
}
