package plog

import (
	"os"
	"fmt"
)

type broadcastLogger struct {
	receiver []Logger
	logger
}

// NewBroadcastLogger creates a new logger, whe the log messages
// are forwarded to the given loggers.
func NewBroadcastLogger(others ...Logger) *broadcastLogger {
	self := new(broadcastLogger)
	initLogger(self)

	for _, logger := range others {
		self.Add(logger)
	}

	return self
}

func (self *broadcastLogger) SetLevel(level LogLevel) {
	for _, logger := range self.receiver {
		logger.SetLevel(level)
	}
}

func (self *broadcastLogger) Add(other Logger) *broadcastLogger {
	self.receiver = append(self.receiver, other)
	return self
}

// Reset removes all attached logger instances
func (self *broadcastLogger) Reset() *broadcastLogger {
	self.receiver = make([]Logger, 0)
	return self
}


func (self *broadcastLogger) log(msg LogMessage) {
	if !dropUnhandledMessages && len(self.receiver) == 0 {
		fmt.Fprintf(os.Stderr, "WARNING - unhandled message: %s\n", msg.Message)
	}

	for _, logger := range self.receiver {
		logger.(logImpl).log(msg)
	}
}
