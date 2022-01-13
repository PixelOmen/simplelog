package main

import (
	"os"
	"testapp/src/levelslog"
)

func main() {
	logfile, _ := os.OpenFile("testlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	mylogger := levelslog.New("Mylog", logfile, true)
	mylogger.SetLevel(levelslog.INFO)
	mylogger.Debug("This is a test")
	mylogger.Info("This is a test")
	mylogger.Warning("This is a test")
	mylogger.Err("This is a test")
	anotherlogger := levelslog.Get("Mylog")
	anotherlogger.Info("From another pointer")
}
