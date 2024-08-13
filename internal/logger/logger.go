package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

var (
	Debug bool
	Info  *Logger
	Warn  *Logger
	Error *Logger
)

func init() {
	Debug = os.Getenv("DEBUG") == "true"

	Info = &Logger{log.New(os.Stdout, "\033[0;34mINFO: \033[0m", log.Ldate|log.Ltime|log.Lshortfile)}
	Warn = &Logger{log.New(os.Stdout, "\033[0;33mWARN: \033[0m", log.Ldate|log.Ltime|log.Lshortfile)}
	Error = &Logger{log.New(os.Stderr, "\033[0;31mERROR: \033[0m", log.Ldate|log.Ltime|log.Lshortfile)}

	if Debug {
		Info.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
		Warn.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
		Error.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	}
}
