package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	ModeDevelopment int8 = iota
	ModeProduct
)

type Logger struct {
	std   *log.Logger
	mode  int8
	depth int
}

func (l *Logger) SetMode(mode int8) {
	l.mode = mode
}

func (l *Logger) SetDepth(depth int) {
	l.depth = depth
}

func (l *Logger) SetOutput(w io.Writer) {
	l.std.SetOutput(w)
}

func (l *Logger) SetFlag(flag int) {
	l.std.SetFlags(flag)
}

func (l *Logger) Output(level, mesaage string) {
	_ = l.std.Output(l.depth, "["+level+"] "+mesaage)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.mode == ModeDevelopment {
		l.Output(lvDebug, fmt.Sprint(v...))
	}
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.mode == ModeDevelopment {
		l.Output(lvDebug, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(lvInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(lvInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.Output(lvWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(lvWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(lvError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(lvError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(lvFatal, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(lvFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Output(lvPanic, s)
	panic(s)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(lvPanic, s)
	panic(s)
}

func (l *Logger) Catch(err error, v ...interface{}) {
	if err != nil {
		l.Error(v, err)
	}
}
