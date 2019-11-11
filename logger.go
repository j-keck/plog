package plog

import (
	"strings"
	"fmt"
)


// logger provides the basic functions of every logger.
//
// Each concrete logger implements his own 'log(LogMessage)' function.
// The basic functions defined on this struct calls the underlying
// 'log(LogMessage)' (abstracted over the 'logImpl' interface) function.
type logger struct {
	level       LogLevel
	logImpl
}
type logImpl interface {
	log(LogMessage)
}

// log level functions
func (self *logger) SetLevel(level LogLevel) {
	self.level = level
}
func (self *logger) IsTraceEnabled() bool {
	return self.level <= Trace
}
func (self *logger) IsDebugEnabled() bool {
	return self.level <= Debug
}
func (self *logger) IsInfoEnabled() bool {
	return self.level <= Info
}
func (self *logger) IsWarnEnabled() bool {
	return self.level <= Warn
}
func (self *logger) IsErrorEnabled() bool {
	return self.level <= Error
}

// logging functions
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
