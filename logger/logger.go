package logger

import (
	"log"
	"os"
	"sync"
)

var (
	loggerInstanse *log.Logger
	once           sync.Once
)

func getLogger() *log.Logger {
	filename := "/home/eyvaz/Documents/lsp_log.txt"

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("oh some fuck happened")
	}

	return log.New(logfile, "[educational lsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLogger() *log.Logger {
	once.Do(func() {
		loggerInstanse = getLogger()
	})
	return loggerInstanse
}
