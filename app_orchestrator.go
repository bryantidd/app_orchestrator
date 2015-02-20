package main

import (
    "encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK int = 0

	ExitCodeError = 10 + iota
	ExitCodeInterrupt
	ExitCodeParseFlagsError
	ExitCodeParseWaitError
	ExitCodeRunnerError
)

type Configuration struct {
    Application struct {
            Name		string  
            ConsulUri  	string
            SwarmUri	string
            NumberOfContainers uint64
			Service		struct {
				ServiceName				string
				ServiceTags				[]string
				ServiceCheckHttp		string
				ServiceCheckInterval	uint64
			}
			Environment		[]map[string]string
			Volumes			[]string
			Restart			string
			Port			string
			Image			string
        }
}
type APP struct {
	
	outStream, errStream io.Writer
	// stopCh is an internal channel used to trigger a shutdown of the APP.
	stopCh chan struct{}
}

func NewApp(out, err io.Writer) *APP {
	return &APP{
		outStream: out,
		errStream: err,
		stopCh:    make(chan struct{}),
	}
}

func (app *APP) stop() {
	close(app.stopCh)
}

func (app *APP) handleError(err error, status int) int {
	fmt.Fprintf(app.errStream, "App Orchestrator returned errors:\n%s", err)
	return status
}

func (app *APP) Run(args []string) int {

    file, _ := os.Open(args[1])
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  fmt.Println("error:", err)
	}
	fmt.Println(configuration.Application)

    return ExitCodeOK

}



