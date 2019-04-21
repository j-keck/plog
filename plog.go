package plog

import (
	"fmt"
	"strings"
)

type Logger interface {
	SetLevel(LogLevel)
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



func (self *logger) Trace(xs ...interface{}) {
	self.log(newLogMessage(Trace, strings.Trim(fmt.Sprintln(xs...), "\n")))
}
func (self *logger) Tracef(format string, xs ...interface{}) {
	self.log(newLogMessage(Trace, fmt.Sprintf(format, xs...)))
}
func (self *logger) Debug(xs ...interface{}) {
	self.log(newLogMessage(Debug, strings.Trim(fmt.Sprintln(xs...), "\n")))
}
func (self *logger) Debugf(format string, xs ...interface{}) {
	self.log(newLogMessage(Debug, fmt.Sprintf(format, xs...)))
}
func (self *logger) Info(xs ...interface{}) {
	self.log(newLogMessage(Info, strings.Trim(fmt.Sprintln(xs...), "\n")))
}
func (self *logger) Infof(format string, xs ...interface{}) {
	self.log(newLogMessage(Info, fmt.Sprintf(format, xs...)))
}
func (self *logger) Warn(xs ...interface{}) {
	self.log(newLogMessage(Warn, strings.Trim(fmt.Sprintln(xs...), "\n")))
}
func (self *logger) Warnf(format string, xs ...interface{}) {
	self.log(newLogMessage(Warn, fmt.Sprintf(format, xs...)))
}
func (self *logger) Error(xs ...interface{}) {
	self.log(newLogMessage(Error, strings.Trim(fmt.Sprintln(xs...), "\n")))
}
func (self *logger) Errorf(format string, xs ...interface{}) {
	self.log(newLogMessage(Error, fmt.Sprintf(format, xs...)))
}
func (self *logger) Fatal(xs ...interface{}) {
	self.log(newLogMessage(Fatal, strings.Trim(fmt.Sprintln(xs...), "\n")))
}
func (self *logger) Fatalf(format string, xs ...interface{}) {
	self.log(newLogMessage(Fatal, fmt.Sprintf(format, xs...)))
}
