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

var gLogger *Logger

func New(logName string, logPath string) *Logger {
	fileName := fmt.Sprintf("%s-%s.log", logName, time.Now().Format("2006-01-02"))
	filePath := path.Join(logPath, fileName)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, "", log.Ldate|log.Lmicroseconds|log.Llongfile)
	return &Logger{
		mLogger: logger,
		file:    file,
	}
}

func (logger *Logger) initLog(level string, msg string) {
	logger.mLogger.Output(2, fmt.Sprintf("[%s]%s", level, msg))
}
func (logger *Logger) Debug(msg string) {
	logger.initLog("DEBUG", msg)
}
func (logger *Logger) Info(msg string) {
	logger.initLog("INFO", msg)
}
func (logger *Logger) Notice(msg string) {
	logger.initLog("NOTICE", msg)
}
func (logger *Logger) Warn(msg string) {
	logger.initLog("WARN", msg)
}
func (logger *Logger) Error(msg string) {
	logger.initLog("ERROR", msg)
}
func (logger *Logger) Close() {
	logger.file.Close()
}
func Debug(msg string) {
	gLogger.Debug(msg)
}

func Info(msg string) {
	gLogger.Info(msg)
}

func Notice(msg string) {
	gLogger.Notice(msg)
}

func Warn(msg string) {
	gLogger.Warn(msg)
}

func Error(msg string) {
	gLogger.Error(msg)
}

func Close() {
	gLogger.Close()
}
