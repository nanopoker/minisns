package logger

import (
	"log"
	"os"
)

var (
	LogFile       *os.File
	TraceLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	_, err := os.Stat("log")
	if os.IsNotExist(err) {
		os.Mkdir("log", 0777)
	}

	LogFile, err = os.OpenFile("log/minisns.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// separate log levels
	TraceLogger = log.New(LogFile,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger = log.New(LogFile,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	WarningLogger = log.New(LogFile,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLogger = log.New(LogFile,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)
}

func Info(message ...interface{}) {
	InfoLogger.Println(message)
}

func Trace(message ...interface{}) {
	TraceLogger.Println(message)
}

func Warning(message ...interface{}) {
	WarningLogger.Println(message)
}

func Error(message ...interface{}) {
	ErrorLogger.Println(message)
}
