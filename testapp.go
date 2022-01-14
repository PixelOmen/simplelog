package main

import (
	"os"

	"github.com/pixelomen/simplelog"
)

func main() {
	logfile, _ := os.OpenFile("testlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	mylogger := simplelog.New("Mylog", logfile, true)
	mylogger.SetLevel(simplelog.INFO)
	mylogger.Debug("This is a test")
	mylogger.Info("This is a test")
	mylogger.Warning("This is a test")
	mylogger.Err("This is a test")
	anotherlogger := simplelog.Get("Mylog")
	anotherlogger.Info("From another pointer")
}
