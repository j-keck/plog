package plog

import (
	"flag"
	"strconv"
)

// debug level logging
type flagDebugValue LogLevel

func FlagDebugVar(p *LogLevel, name string, usage string) {
	flag.CommandLine.Var((*flagDebugValue)(p), name, usage)
}

func (self *flagDebugValue) Set(str string) error {
	if flagIsTrue, err := strconv.ParseBool(str); err != nil {
		return err
	} else {
		if flagIsTrue && (*self > flagDebugValue(Debug)) {
			*self = flagDebugValue(Debug)
		}
	}
	return nil
}

func (self *flagDebugValue) String() string {
	return "Debug"
}

func (self *flagDebugValue) IsBoolFlag() bool {
	return true
}


// trace level logging
type flagTraceValue LogLevel

func FlagTraceVar(p *LogLevel, name string, usage string) {
	flag.CommandLine.Var((*flagTraceValue)(p), name, usage)
}

func (self *flagTraceValue) Set(str string) error {
	if flagIsTrue, err := strconv.ParseBool(str); err != nil {
		return err
	} else {
		if flagIsTrue {
			*self = flagTraceValue(Trace)
		}
	}
	return nil
}

func (self *flagTraceValue) String() string {
	return "Trace"
}

func (self *flagTraceValue) IsBoolFlag() bool {
	return true
}
