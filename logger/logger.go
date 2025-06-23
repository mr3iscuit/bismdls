package logger

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sync"
)

var (
	loggerInstanse *log.Logger
	once           sync.Once
)

func getLogger() *log.Logger {
	filePath := "Documents/lsp_log.txt"

	usr, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("could not get the current user: %v", err))
	}

	homedir := usr.HomeDir

	logfile, err := os.OpenFile(filepath.Join(homedir, filePath), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("could not open the log file")
	}

	return log.New(logfile, "[educational lsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLogger() *log.Logger {
	once.Do(func() {
		loggerInstanse = getLogger()
	})
	return loggerInstanse
}
