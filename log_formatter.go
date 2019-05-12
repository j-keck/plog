package plog

import (
	"fmt"
	"strings"
)

// LogFormatter formats a LogMessage
type LogFormatter struct {
	prefix      string
	separator   string
	suffix      string
	formatters  []Formatter
}

// NewLogFormatter create a LogFormatter with the given separator string
// and the given column formatters.
func NewLogFormatter(sep string, formatters ...Formatter) LogFormatter {
	return LogFormatter{"", sep, "", formatters}
}

// SetLogPrefix sets a prefix which will be prepended to each line.
func (self *LogFormatter) SetLogPrefix(prefix string) *LogFormatter {
	self.prefix = prefix
	return self
}

// SetLogSuffix sets a prefix which will be appended to each line.
func (self *LogFormatter) SetLogSuffix(suffix string) *LogFormatter {
	self.suffix = suffix
	return self
}

// AddLogFormatter adds a column formatter
func (self *LogFormatter) AddLogFormatter(formatter Formatter) *LogFormatter {
	self.formatters = append(self.formatters, formatter)
	return self
}

// Format the give log message
func (self *LogFormatter) Format(msg LogMessage) string {
	var builder strings.Builder

	for _, formatter := range self.formatters {
		if builder.Len() > 0 {
			builder.WriteString(self.separator)
		}
		builder.WriteString(formatter.Format(&msg))
	}

	return self.prefix + builder.String() + self.suffix
}




// Predefined formatters
const (
	Level             = LevelFmt("%5s")
	Timestamp         = TimestampFmt("Jan _2 15:04:05")
	TimestampMillis   = TimestampFmt("Jan _2 15:04:05.000")
	TimestampUnixDate = TimestampFmt("Mon Jan _2 15:04:05 MST 2006")
	Location          = LocationFmt("%15s:%-3d")
	File              = FileFmt("%15s")
	Line              = LineFmt("%-3d")
	Message           = MessageFmt("%s")
)

type Formatter interface {
	Format(*LogMessage) string
}

// LevelFmt describes how the LogLevel are formatted.
//
// Typical format is "%5s".
type LevelFmt string

func (self LevelFmt) Format(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.Level)
}

// TimestampFmt describes how the timestamps are formatted.
//
// For a description of valid format parameters, see:
//   https://golang.org/pkg/time/#Time.Format
type TimestampFmt string

func (self TimestampFmt) Format(msg *LogMessage) string {
	return msg.Timestamp.Format(string(self))
}

// FileFmt describes how the log-caller filename are formatted.
//
// Typical format is "%15s".
type FileFmt string

func (self FileFmt) Format(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.File)
}

// LineFmt describes how the log-caller source line are formatted.
//
// Typical format is "%-3d"
type LineFmt string

func (self LineFmt) Format(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.Line)
}

// LocationFmt describes how the log-caller location are formatted.
//
// The location combines the log-caller filename and source line.
type LocationFmt string

func (self LocationFmt) Format(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.File, msg.Line)
}

// MessageFmt describes how the log-message are formatted.
//
// Typical format is "%s".
type MessageFmt string

func (self MessageFmt) Format(msg *LogMessage) string {
	return fmt.Sprintf(string(self), msg.Message)
}
