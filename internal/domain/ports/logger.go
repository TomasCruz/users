package ports

type LogLvl int

const (
	DEBUG_LOG_LEVEL LogLvl = iota + 1
	INFO_LOG_LEVEL
	WARN_LOG_LEVEL
	ERROR_LOG_LEVEL
	FATAL_LOG_LEVEL
)

type Logger interface {
	Debug(err error, msg string)
	Info(err error, msg string)
	Warn(err error, msg string)
	Error(err error, msg string)
	Fatal(err error, msg string)
}

func LogLvlToString(level LogLvl) string {
	switch level {
	case DEBUG_LOG_LEVEL:
		return "DEBUG"
	case INFO_LOG_LEVEL:
		return "INFO"
	case WARN_LOG_LEVEL:
		return "WARN"
	case ERROR_LOG_LEVEL:
		return "ERROR"
	case FATAL_LOG_LEVEL:
		return "FATAL"
	}

	return ""
}

func StringToLogLvl(lvl string) LogLvl {
	switch lvl {
	case "DEBUG":
		return DEBUG_LOG_LEVEL
	case "INFO":
		return INFO_LOG_LEVEL
	case "WARN":
		return WARN_LOG_LEVEL
	case "ERROR":
		return ERROR_LOG_LEVEL
	case "FATAL":
		return FATAL_LOG_LEVEL
	}

	return DEBUG_LOG_LEVEL
}
