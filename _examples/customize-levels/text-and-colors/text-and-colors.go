package main

import (
	"github.com/iami317/logx"
)

func main() {

	// First argument is the raw text for outputs
	// second argument is the color code
	// and the last, variadic argument can be any `kataras/pio.RichOption`, e.g. pio.Background, pio.Underline.

	// Default is "[ERRO]"
	logx.ErrorText("|ERROR|", 31)
	// Default is "[WARN]"
	logx.WarnText("|WARN|", 32)
	// Default is "[INFO]"
	logx.InfoText("|INFO|", 34)
	// Default is "[DBUG]"
	logx.DebugText("|DEBUG|", 33)

	// Business as usual...
	logx.SetLevel("debug")

	logx.Println("This is a raw message, no levels, no colors.")
	logx.Info("This is an info message, with colors (if the output is terminal)")
	logx.Warn("This is a warning message")
	logx.Error("This is an error message")
	logx.Debug("This is a debug message")
}
