package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/TomasCruz/users/internal/domain/entities"
	"github.com/pkg/errors"
)

type printLogger struct {
	minLevel entities.LogLvl
}

func New(minLevel entities.LogLvl) entities.Logger {
	return printLogger{minLevel: minLevel}
}

func (p printLogger) Debug(err error, msg string) {
	if entities.DEBUG_LOG_LEVEL >= p.minLevel {
		s := p.logToString(entities.DEBUG_LOG_LEVEL, err, msg)
		fmt.Print(s)
	}
}

func (p printLogger) Info(err error, msg string) {
	if entities.INFO_LOG_LEVEL >= p.minLevel {
		s := p.logToString(entities.INFO_LOG_LEVEL, err, msg)
		fmt.Print(s)
	}
}

func (p printLogger) Warn(err error, msg string) {
	if entities.WARN_LOG_LEVEL >= p.minLevel {
		s := p.logToString(entities.WARN_LOG_LEVEL, err, msg)
		fmt.Print(s)
	}
}

func (p printLogger) Error(err error, msg string) {
	if entities.ERROR_LOG_LEVEL >= p.minLevel {
		s := p.logToString(entities.ERROR_LOG_LEVEL, err, msg)
		fmt.Print(s)
	}
}

func (p printLogger) Fatal(err error, msg string) {
	s := p.logToString(entities.FATAL_LOG_LEVEL, err, msg)
	fmt.Print(s)
	os.Exit(1)
}

func (p printLogger) logToString(level entities.LogLvl, err error, msg string) string {
	var sb strings.Builder

	sb.WriteString(entities.LogLvlToString(level))
	sb.WriteRune(' ')
	sb.WriteString(time.Now().Format("2006-01-02 15:04:05.000"))
	sb.WriteRune('\n')

	if msg != "" {
		sb.WriteString(msg)
	}

	if err != nil {
		sb.WriteString(":\n\t")
		sb.WriteString(err.Error())
		sb.WriteRune('\n')
		sb.WriteString(errStackToString(err))
	}

	return sb.String()
}

func errStackToString(err error) string {
	var sb strings.Builder

	calls := funcCalls(err)
	for _, fc := range calls {
		sb.WriteString(fmt.Sprintf("\t%s %d -\t%s\n", fc.fileName, fc.line, fc.funcName))
	}

	return sb.String()
}

type functionCall struct {
	filePath string
	fileName string
	funcName string
	line     int
}

func funcCalls(err error) []functionCall {
	var calls []functionCall
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
			calls = append(calls, fc)
		}
	}

	return calls
}

func fileName(name string) string {
	i := strings.LastIndex(name, "/")
	return name[i+1:]
}
