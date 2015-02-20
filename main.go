package main

import (
	"os"
)

// Name is the exported name of this application.
const Name = "app_orchestrator"

// Version is the current version of this application.
const Version = "0.0.1.bnt"

func main() {
	app := NewApp(os.Stdout, os.Stderr)
	os.Exit(app.Run(os.Args))
}