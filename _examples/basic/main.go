package main

import (
	"github.com/iami317/logx"
)

func main() {
	// Default Output is `os.Stdout`,
	// but you can change it:
	// golog.SetOutput(os.Stderr)

	// Time Format defaults to: "2006/01/02 15:04"
	// you can change it to something else or disable it with:
	log := logx.New()
	log.MessageColor = true
	log.TimeColor = true
	logx.SetTimeFormat("")

	// Level defaults to "info",
	// but you can change it:
	//logx.SetLevel("debug")

	log.Println("This is a raw message, no levels, no colors.")
	logx.Verbose("This is a verbose message, it will be printed if logger's Level is <= VerboseLevel")
	logx.Debug("This is a debug message")
	log.Info("This is an info message, with colors (if the output is terminal)")
	log.Warn("This is a warning message")
	logx.Error("This is an error message")
	logx.Silentf("This is a silent message, it will not be printed")
	logx.Fatal("Fatal will exit no matter what, but it will also print the log message if logger's Level is >= FatalLevel")
}
