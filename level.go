package plog

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Info
	Note
	Warn
	Error
	Fatal
)

func (l LogLevel) String() string {
	switch l {
	case Trace:
		return "TRACE"
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Note:
		return "NOTE"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return "<IMPOSSIBLE>"
}
