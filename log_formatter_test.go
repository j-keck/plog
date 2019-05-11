package plog

import "testing"
import "time"


func TestLogFormatter(t *testing.T) {
	fmt := NewLogFormatter(" | ",
		Level,
		TimestampUnixDate,
		Location,
		Message)
	fmt.SetLogPrefix("[").SetLogSuffix("]")

	expected := "[ INFO | Thu Jan  1 01:00:00 CET 1970 |        filename:33  | Test]"
	actual := fmt.Format(LogMessage{Info, time.Unix(0, 0), "filename", 33, "Test"})
	if expected != actual {
		t.Errorf("expected != actual\n exp: '%s'\n act: '%s'", expected, actual)
	}
}


func TestFormatter(t *testing.T) {
	type test struct {
		formatter Formatter
		expected  string
	}

	tests := []test{
		// provided
		test{ Level, " INFO"},
		test{ Timestamp, "Jan  1 01:00:00" },
		test{ TimestampMillis, "Jan  1 01:00:00.000" },
		test{ TimestampUnixDate, "Thu Jan  1 01:00:00 CET 1970" },
		test{ File, "       filename" },
		test{ Line, "33 " },
		test{ Location, "       filename:33 "},
		test{ Message, "Test" },

		// custom
		test{ LevelFmt("(%s)"), "(INFO)" },
		test{ MessageFmt("msg: %s"), "msg: Test" },
	}

	msg := LogMessage{Info, time.Unix(0, 0), "filename", 33, "Test"}
	for _, test := range tests {
		actual := test.formatter.Format(&msg)
		if test.expected != actual {
			t.Errorf("expected != actual\n exp: '%s'\n act: '%s'", test.expected, actual)
		}
	}

}

