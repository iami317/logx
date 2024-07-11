package main

import (
	"time"

	"github.com/iami317/logx"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

/*
	go get github.com/lestrrat-go/file-rotatelogs
*/

func main() {
	// Read more at: https://github.com/lestrrat-go/file-rotatelogs#synopsis
	pathToAccessLog := "./access_log.%Y%m%d%H%M"
	w, err := rotatelogs.New(
		pathToAccessLog,
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour))
	if err != nil {
		logx.Fatal(err)
	}
	defer w.Close()

	logx.SetOutput(w)

	logx.Println("A Log entry")
}
