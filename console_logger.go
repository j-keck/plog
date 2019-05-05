package plog

import (
	"fmt"
	"os"
	"io"
)

type consoleLogger struct {
	logger
	stdout io.Writer
	stderr io.Writer
}

// NewConsoleLogger creates a new logger, where the log messages
// are emitted to the console
func NewConsoleLogger() *consoleLogger {
	self := new(consoleLogger)
	initLogger(self)
	self.SetStdout(os.Stdout)
	self.SetStderr(os.Stderr)
	return self
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

	ts := msg.Timestamp.Format("02.01 15:04:05.000")
	logStr := fmt.Sprintf("%s | %s | %20s:%-3d | %s", msg.Level, ts, msg.File, msg.Line, msg.Message)

	fmt.Fprintln(out, logStr)
}
