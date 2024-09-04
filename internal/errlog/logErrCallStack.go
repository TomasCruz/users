package errlog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type LogLvl int

const (
	DEBUG LogLvl = iota + 1
	INFO
	WARN
	ERROR
	FATAL
)

func Debug(err error, msg string) {
	s := ToString(DEBUG, err, msg)
	fmt.Print(s)
}

func Info(err error, msg string) {
	s := ToString(INFO, err, msg)
	fmt.Print(s)
}

func Warn(err error, msg string) {
	s := ToString(WARN, err, msg)
	fmt.Print(s)
}

func Error(err error, msg string) {
	s := ToString(ERROR, err, msg)
	fmt.Print(s)
}

func Fatal(err error, msg string) {
	s := ToString(FATAL, err, msg)
	fmt.Print(s)
	os.Exit(1)
}

func ToString(level LogLvl, err error, msg string) string {
	var sb strings.Builder

	sb.WriteString(logLvlToString(level))
	sb.WriteRune(' ')
	sb.WriteString(time.Now().Format("2006-01-02 15:04:05.000"))
	sb.WriteRune(' ')

	if msg != "" {
		sb.WriteString(fmt.Sprintf("%s:\n  ", msg))
	}

	sb.WriteString(err.Error())
	sb.WriteRune('\n')
	sb.WriteString(errStackToString(err))

	return sb.String()
}

func errStackToString(err error) string {
	var sb strings.Builder

	functionCalls := errlog(err)
	for _, fc := range functionCalls {
		sb.WriteString(fmt.Sprintf("\t%s %d -\t%s\n", fc.fileName, fc.line, fc.funcName))
	}

	return sb.String()
}

func logLvlToString(level LogLvl) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}

	return ""
}

type functionCall struct {
	filePath string
	fileName string
	funcName string
	line     int
}

func errlog(err error) (functionCalls []functionCall) {
	if err, ok := err.(interface{ StackTrace() errors.StackTrace }); ok {
		for _, f := range err.StackTrace() {
			pc := uintptr(f) - 1
			fn := runtime.FuncForPC(pc)
			file, line := fn.FileLine(pc)
			fc := functionCall{
				filePath: file,
				fileName: fileName(file),
				funcName: fn.Name(),
				line:     line,
			}
			functionCalls = append(functionCalls, fc)
		}
	}

	return
}

func fileName(name string) string {
	i := strings.LastIndex(name, "/")
	return name[i+1:]
}
