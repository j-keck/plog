package plog

import (
	"fmt"
	"os"
)

type consoleLogger struct {
	logger
}

// NewConsoleLogger creates a new logger, where the log messages
// are emitted to the console
func NewConsoleLogger() Logger {
	self := new(consoleLogger)
	self.SetLevel(Info)
	self.SetStdout(os.Stdout)
	self.SetStderr(os.Stderr)

	// FIXME: remove this bad hack.
	self.logImpl = self

	return self
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
