package plog

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type Logger interface {
	SetLevel(LogLevel)
	SetStdout(io.Writer)
	SetStderr(io.Writer)
	Subscribe(int) <-chan LogMessage
	// FIXME: return errors? / counters?
	WaitForSubscribers(time.Duration)
	Trace(...interface{})
	Tracef(string, ...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

type logger struct {
	level       LogLevel
	stdout      io.Writer
	stderr      io.Writer
	subscribers []chan LogMessage
	logImpl
}

// FIXME: remove this bad hack
// to call the 'log' method from a "subclass".
// the "subclass" needs to set the 'logImpl' per 'self.logImpl = self'!
type logImpl interface {
	log(LogMessage)
}

func (self *logger) SetLevel(level LogLevel) {
	self.level = level
}

// SetStdout to overwrite 'stdout'
func (self *logger) SetStdout(w io.Writer) {
	self.stdout = w
}

// SetStderr to overwrite 'stderr'
func (self *logger) SetStderr(w io.Writer) {
	self.stderr = w
}

// Subscribe to log messages
func (self *logger) Subscribe(bufferSize int) <-chan LogMessage {
	sub := make(chan LogMessage, bufferSize)
	self.subscribers = append(self.subscribers, sub)
	return sub
}

func (self logger) WaitForSubscribers(timeout time.Duration) {
	for _, c := range self.subscribers {
		const attempts = 10
		for i := 0; ; i++ {
			if i >= attempts {
				timeoutMsg := fmt.Sprintf("plog.WaitForSubscribers timeout after %s - ignore Subscriber", timeout.String())
				logMsg := newLogMessage(Warn, timeoutMsg)
				fmt.Fprintln(self.stderr, logMsg.String())
				break
			}

			if len(c) == 0 {
				break
			}

			time.Sleep(timeout / attempts)
		}
	}
}

func (self *logger) Trace(xs ...interface{}) {
	msg := newLogMessage(Trace, strings.Trim(fmt.Sprintln(xs...), "\n"))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Tracef(format string, xs ...interface{}) {
	msg := newLogMessage(Trace, fmt.Sprintf(format, xs...))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Debug(xs ...interface{}) {
	msg := newLogMessage(Debug, strings.Trim(fmt.Sprintln(xs...), "\n"))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Debugf(format string, xs ...interface{}) {
	msg := newLogMessage(Debug, fmt.Sprintf(format, xs...))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Info(xs ...interface{}) {
	msg := newLogMessage(Info, strings.Trim(fmt.Sprintln(xs...), "\n"))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Infof(format string, xs ...interface{}) {
	msg := newLogMessage(Info, fmt.Sprintf(format, xs...))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Warn(xs ...interface{}) {
	msg := newLogMessage(Warn, strings.Trim(fmt.Sprintln(xs...), "\n"))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Warnf(format string, xs ...interface{}) {
	msg := newLogMessage(Warn, fmt.Sprintf(format, xs...))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Error(xs ...interface{}) {
	msg := newLogMessage(Error, strings.Trim(fmt.Sprintln(xs...), "\n"))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Errorf(format string, xs ...interface{}) {
	msg := newLogMessage(Error, fmt.Sprintf(format, xs...))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Fatal(xs ...interface{}) {
	msg := newLogMessage(Fatal, strings.Trim(fmt.Sprintln(xs...), "\n"))
	self.log(msg)
	self.notifySubscribers(msg)
}
func (self *logger) Fatalf(format string, xs ...interface{}) {
	msg := newLogMessage(Fatal, fmt.Sprintf(format, xs...))
	self.log(msg)
	self.notifySubscribers(msg)
}

func (self *logger) notifySubscribers(msg LogMessage) {
	if msg.Level < self.level {
		return
	}

	if len(self.subscribers) > 0 {
		for _, sub := range self.subscribers {
			if len(sub) < cap(sub) {
				sub <- msg
			} else {
				overflowMessage := newLogMessage(Warn, "channel full - discard msg: "+msg.String())
				fmt.Fprintln(self.stderr, overflowMessage.String())
			}
		}
	}
}
