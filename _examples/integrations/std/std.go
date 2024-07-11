package main

import (
	"log"
	"os"

	"github.com/iami317/logx"
)

// simulate a log.Logger preparation:
var myLogger = log.New(os.Stdout, "", 0)

func main() {
	logx.SetLevel("error")
	logx.InstallStd(myLogger)

	logx.Debug(`this debug message will not be shown,
	because the golog level is ErrorLevel`)

	logx.Error("this error message will be visible the only visible")

	logx.Warn("this info message will not be visible")
}
