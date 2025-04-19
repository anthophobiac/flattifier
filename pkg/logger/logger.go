package logger

import (
	"github.com/fatih/color"
)

var (
	infoColor    = color.New(color.FgGreen).SprintFunc()
	warningColor = color.New(color.FgYellow).SprintFunc()
	errorColor   = color.New(color.FgRed).SprintFunc()
	debugColor   = color.New(color.FgCyan).SprintFunc()
	prefixColor  = color.New(color.Bold).SprintFunc()
)

func Info(msg string) {
	println(prefixColor("[INFO]"), infoColor(msg))
}

func Warning(msg string) {
	println(prefixColor("[WARN]"), warningColor(msg))
}

func Error(msg string) {
	println(prefixColor("[ERROR]"), errorColor(msg))
}

func Debug(msg string) {
	println(prefixColor("[DEBUG]"), debugColor(msg))
}
