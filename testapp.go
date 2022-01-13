package main

import (
	"os"
	"testapp/src/loglevel"
)

func main() {
	mylogger := loglevel.New(os.Stdout, true)
	mylogger.SetLevel(loglevel.INFO)
	mylogger.Debug("This is a test")
	mylogger.Info("This is a test")
	mylogger.Warning("This is a test")
	mylogger.Err("This is a test")
}
