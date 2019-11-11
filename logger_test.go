package plog

import (
	"testing"
)

func TestCheckLevel(t *testing.T) {
	log := NewStreamLogger()

	// error
	log.SetLevel(Error)
	if !log.IsErrorEnabled() {
		t.Error("error was not enabled")
	}
	if log.IsWarnEnabled() ||
		log.IsInfoEnabled() ||
		log.IsDebugEnabled() ||
		log.IsTraceEnabled() {
		t.Error("wrong level active")
	}

	// warn
	log.SetLevel(Warn)
	if !(log.IsErrorEnabled() &&
		log.IsWarnEnabled()) {
		t.Error("warn was not enabled")
	}
	if log.IsInfoEnabled() ||
		log.IsDebugEnabled() ||
		log.IsTraceEnabled() {
		t.Error("wrong level active")
	}

	// info
	log.SetLevel(Info)
	if !(log.IsErrorEnabled() &&
		log.IsWarnEnabled() &&
		log.IsInfoEnabled()) {
		t.Error("info was not enabled")
	}
	if log.IsDebugEnabled() ||
		log.IsTraceEnabled() {
		t.Error("wrong level active")
	}

	// debug
	log.SetLevel(Debug)
	if !(log.IsErrorEnabled() &&
		log.IsWarnEnabled() &&
		log.IsInfoEnabled() &&
		log.IsDebugEnabled()) {
		t.Error("info was not enabled")
	}
	if log.IsTraceEnabled() {
		t.Error("wrong level active")
	}

	// trace
	log.SetLevel(Trace)
	if !(log.IsErrorEnabled() &&
		log.IsWarnEnabled() &&
		log.IsInfoEnabled() &&
		log.IsDebugEnabled() &&
		log.IsTraceEnabled()) {
		t.Error("trace  was not enabled")
	}
}
