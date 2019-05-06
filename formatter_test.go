package plog

import "testing"

func TestLogFormatter(t *testing.T) {
	fmt := NewLogFormatter(" - ", Level("level: %s"), Message("msg: %s"))
	fmt.SetLogPrefix("(").SetLogSuffix(")")

	expected := "(level: INFO - msg: Test)"
	actual := fmt.Format(newLogMessage(Info, "Test"))
	if expected != actual {
		t.Errorf("%s != %s", expected, actual)
	}
}
