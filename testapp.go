package main

import (
	"os"

	"github.com/pixelomen/simplelog"
)

func main() {
	logfile, _ := os.OpenFile("testlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	mylogger := simplelog.New("Mylog", logfile, true)
	mylogger.SetLevel(simplelog.DEBUG)
	mylogger.Debug("First msg")
	mylogger.Info("First msg")
	mylogger.Warning("First msg")
	mylogger.Err("First msg")

	anotherPointer := simplelog.Get("Mylog")
	anotherPointer.Info("Another pointer")
	anotherPointer.Info("--------------")

	anotherPointer.SetLevel(simplelog.WARNING)
	anotherPointer.Debug("Warnings and higher")
	anotherPointer.Info("Warnings and higher")
	anotherPointer.Warning("Warnings and higher")
	anotherPointer.Err("Warnings and higher")
	stdoutLogger := simplelog.New("naked logger", os.Stdout, false, 0)
	stdoutLogger.Info("Just a prefix")
}
