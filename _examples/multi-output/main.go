package main

import (
	"os"
	"time"

	"github.com/iami317/logx"
)

func main() {
	f := newLogFile()
	defer f.Close()

	logx.AddOutput(f)

	logx.Println("This is a raw message, no levels, no colors.")
	logx.Info("This is an info message, without colors because one of the Outputs is not a terminal-based")
	logx.Warn("This is a warning message")
	logx.Error("This is an error message")
}

// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}
