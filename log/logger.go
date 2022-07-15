package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
}

func Error(a ...any) {
	_, file, line, _ := runtime.Caller(1)

	errorLogger.Printf("%s:%s %s", file, strconv.Itoa(line), a)
}

func Warning(a ...any) {
	_, file, line, _ := runtime.Caller(1)

	errorLogger.Printf("%s:%s %s", file, strconv.Itoa(line), a)
}

func Info(a ...any) {
	_, file, line, _ := runtime.Caller(1)

	infoLogger.Printf("%s:%s %s", file, strconv.Itoa(line), a)
}

func Infof(format string, args ...any) {
	_, file, line, _ := runtime.Caller(1)

	message := fmt.Sprintf(format, args)

	infoLogger.Printf("%s:%s %s", file, strconv.Itoa(line), message)
}
