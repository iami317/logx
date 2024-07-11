package main

import "github.com/iami317/logx"

func main() {
	logx.SetLevel("debug")

	logx.SetFormat("json", "    ") // < --
	// To register a custom formatter:
	// golog.RegisterFormatter(golog.Formatter...)
	/* Example Output:
	{
	    "timestamp": 1591423477,
	    "level": "debug",
	    "message": "This is a message with data (debug prints the stacktrace too)",
	    "fields": {
	        "username": "kataras"
	    },
	    "stacktrace": [
	        {
	            "function": "main.main",
	            "source": "C:/mygopath/src/gitee.com/menciis/logx/_examples/customize-output/main.go:29"
	        }
	    ]
	}
	*/
	logx.Debugf("This is a %s with data (debug prints the stacktrace too)", "message", logx.Fields{
		"username": "kataras",
	}) // If more than one golog.Fields passed, then they are merged into a single map.

	/* Example Output:
	{
	    "timestamp": 1591423477,
	    "level": "info",
	    "message": "An info message",
	    "fields": {
	        "home": "https://iris-go.com"
	    }
		"stacktrace": [...]
	}
	*/
	logx.Infof("An info message", logx.Fields{"home": "https://iris-go.com"})

	logx.Warnf("Hey, warning here")
	logx.Errorf("Something went wrong!")

	// You can also pass custom structs, like normally you would do.
	type myCustomData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	logx.Fatalf("A fatal error for %s screen!", "home", logx.Fields{"data": myCustomData{
		Username: "kataras",
		Email:    "kataras2006@hotmail.com",
	}})
}

/* Manually, use it for any custom format:
golog.Handle(jsonOutput)

func jsonOutput(l *golog.Log) bool {
	enc := json.NewEncoder(l.Logger.Printer)
	enc.SetIndent("", "    ")
	err := enc.Encode(l)
	return err == nil
}
*/
