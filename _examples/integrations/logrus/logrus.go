package main

import (
	"github.com/iami317/logx"
	"github.com/sirupsen/logrus"
)

func main() {
	// outputOnly()
	full()
}

func full() {
	// simulate a logrus preparation:
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// pass logrus.StandardLogger() to print logs using using the default,
	// package-level logrus' instance of Logger:
	logx.Install(logrus.StandardLogger())

	logx.Debug(`this debug message will not be shown,
	because the logrus level is InfoLevel`)
	logx.Error("this error message will be visible as json")

	// simulate a change of the logrus formatter
	// as you see we have nothing more to change
	// on the golog, it works out of the box,
	// it will be adapt by this change, automatically.
	logrus.SetFormatter(&logrus.TextFormatter{})

	logx.Error("this error message will be visible as text")
	logx.Info("this info message will be visible as text")
}

func outputOnly() {
	logx.SetOutput(logrus.StandardLogger().Out)
	logx.Info(`output only, this will print the same contents
	as golog but using the defined logrus' io.Writer`)

	logx.Error("this error message will be visible as text")
}
