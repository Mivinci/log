package log

import (
	"io"
	"log"
	"os"
)

var defaultLogger = &Logger{
	std:   log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	depth: 4,
}

func SetMode(mode int8) {
	defaultLogger.SetMode(mode)
}

func SetDepth(depth int) {
	defaultLogger.SetDepth(depth)
}

func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

func SetFlag(flag int) {
	defaultLogger.SetFlag(flag)
}

func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v...)
}

func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

func Panic(v ...interface{}) {
	defaultLogger.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	defaultLogger.Panicf(format, v...)
}

func Catch(err error, v ...interface{}) {
	defaultLogger.Catch(err, v)
}
