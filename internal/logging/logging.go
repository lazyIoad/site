package logging

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	PanicLogger   *log.Logger
)

func InitLogging() {
	InfoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	PanicLogger = log.New(os.Stderr, "[PANIC] ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
}
