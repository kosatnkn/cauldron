package log

import (
	"github.com/gookit/color"
)

// Error logs a message as of error type.
func Error(msg string) {
	log("ERROR", msg)
}

// Debug logs a message as of debug type.
func Debug(msg string) {
	log("DEBUG", msg)
}

// Info logs a message as of information type.
func Info(msg string) {
	log("INFO", msg)
}

// Warn logs a message as of warning type.
func Warn(msg string) {
	log("WARN", msg)
}

// Theme logs a message as of theme type.
func Theme(msg string) {
	log("THEME", msg)
}

// Note logs a message as of notification type.
func Note(msg string) {
	log("NOTE", msg)
}

// Default logs a message as of default type.
func Default(msg string) {
	log("", msg)
}

// log logs a message with colour.
func log(level, msg string) {

	switch level {
	case "ERROR":
		color.New(color.FgRed, color.OpBold).Println("ERROR: " + msg)
		break
	case "DEBUG":
		color.Debug.Println(msg)
		break
	case "INFO":
		color.Info.Println(msg)
		break
	case "WARN":
		color.Warn.Println(msg)
		break
	case "THEME":
		color.Note.Println(msg)
		break
	case "NOTE":
		color.Note.Println(msg)
		break
	default:
		color.Println(msg)
		break
	}
}
