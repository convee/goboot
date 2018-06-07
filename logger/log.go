package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type Logger struct {
	mLogger *log.Logger
	file    *os.File
}

var logger *Logger

func Init(logName string, logPath string) {
	fileName := fmt.Sprintf("%s-%s.log", logName, time.Now().Format("2006-01-02"))
	filePath := path.Join(logPath, fileName)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	mLogger := log.New(file, "", log.Ldate|log.Lmicroseconds|log.Llongfile)
	logger = &Logger{
		mLogger: mLogger,
		file:    file,
	}
}
func writeLog(level string, msg string) {
	logger.mLogger.Output(2, fmt.Sprintf("[%s]%s", level, msg))
}
func Close() {
	logger.file.Close()
}
func Debug(msg string) {
	writeLog("DEBUT", msg)
}
func Info(msg string) {
	writeLog("INFO", msg)
}
func Notice(msg string) {
	writeLog("NOTICE", msg)
}
func Warn(msg string) {
	writeLog("WARN", msg)
}
func Error(msg string) {
	writeLog("ERROR", msg)
}
