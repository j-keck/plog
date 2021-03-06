package plog

import (
	"fmt"
	"os"
	"io"
)

type consoleLogger struct {
	logger
    LogFormatter
	stdout     io.Writer
	stderr     io.Writer
}

// NewConsoleLogger creates a new logger, where the log messages
// are emitted to the console
func NewConsoleLogger(separator string, formatters ...Formatter) *consoleLogger {
	self := new(consoleLogger)
	initLogger(self)
	self.LogFormatter = NewLogFormatter(separator, formatters...)
	self.SetStdout(os.Stdout)
	self.SetStderr(os.Stderr)
	return self
}

func NewDefaultConsoleLogger() *consoleLogger {
	return NewConsoleLogger(" | ",
		TimestampUnixDate,
		Level,
		Message,
	)
}

// SetStdout to overwrite 'stdout'
func (self *consoleLogger) SetStdout(w io.Writer) {
	self.stdout = w
}

// SetStderr to overwrite 'stderr'
func (self *consoleLogger) SetStderr(w io.Writer) {
	self.stderr = w
}


func (self *consoleLogger) log(msg LogMessage) {
	if msg.Level < self.level {
		return
	}

	out := self.stdout
	if msg.Level >= Warn {
		out = self.stderr
	}

	fmt.Fprintln(out, self.Format(msg))
}
