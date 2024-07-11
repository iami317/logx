package benchmarks

import (
	"testing"

	"gitee.com/menciis/logx"
)

var nopOutput = logx.NopOutput

func BenchmarkGologPrint(b *testing.B) {
	// logger defaults
	logx.SetOutput(nopOutput)
	logx.SetLevel("debug")
	// disable time formatting because logrus and std doesn't print the time.
	// note that the time is being set-ed to time.Now() inside the golog's Log structure, same for logrus,
	// Therefore we set the time format to empty on golog test in order
	// to acomblish a fair comparison between golog and logrus.
	logx.SetTimeFormat("")

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		printGolog(i)
	}
}

func printGolog(i int) {
	logx.Errorf("[%d] This is an error message", i)
	logx.Warnf("[%d] This is a warning message", i)
	logx.Infof("[%d] This is an info message", i)
	// Debug on golog prints the whole stacktrace, while logrus does not, don't include that.
}
