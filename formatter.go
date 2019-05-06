package plog

import (
	"fmt"
	"strings"
)

// LogFormatter formats a LogMessage
//
//
type LogFormatter struct {
	prefix    string
	separator string
	suffix    string
	fields    []FieldFmt
}

// NewLogFormatter create a LogFormatter with the given sparator string
// and the given field formats.
func NewLogFormatter(sep string, fields ...FieldFmt) LogFormatter  {
	return LogFormatter { "", sep, "", fields }
}

// NewDefaultLogFormatter creates a LogFormatter with a default format.
//
// Example log output:
//
//    INFO | 06.05 10:02:11.033 |              console:8   | startup
//   DEBUG | 06.05 10:02:11.033 |              console:11  | level changed
//
func NewDefaultLogFormatter() LogFormatter {
	return NewLogFormatter(" | ",
		Level("%5s"),
		Timestamp("02.01 15:04:05.000"),
		Location("%20s:%-3d"),
		Message("%s"),
	)
}

func (self *LogFormatter) SetLogPrefix(prefix string) *LogFormatter{
	self.prefix = prefix
	return self
}

func (self *LogFormatter) SetLogSeparator(separator string) *LogFormatter {
	self.separator = separator
	return self
}

func (self *LogFormatter) SetLogFields(fields ...FieldFmt) *LogFormatter {
	self.fields = fields
	return self
}

func (self *LogFormatter) SetLogSuffix(suffix string) *LogFormatter {
	self.suffix = suffix
	return self
}

// Format the give log message
func (self *LogFormatter) Format(msg LogMessage) string {
	var builder strings.Builder

	for _, field := range self.fields {
		if builder.Len() > 0 {
			builder.WriteString(self.separator)
		}
		builder.WriteString(field.fmt(&msg))
	}

	return self.prefix + builder.String() + self.suffix
}


type FieldFmt interface {
	fmt(*LogMessage) string
}


// Level("%5s")
type Level string
func (self Level) fmt(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.Level)
}

// Timestamp("02.01 15:04:05.00")
type Timestamp string
func (self Timestamp) fmt(msg *LogMessage) string {
	return msg.Timestamp.Format(string(self))
}

// File("%s")
type File string
func (self File) fmt(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.File)
}

// Line("%d")
type Line string
func (self Line) fmt(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.Line)
}

// Location("%s:%02d")
type Location string
func (self Location) fmt(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.File, msg.Line)
}

// Message("%s")
type Message string
func (self Message) fmt(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.Message)
}
