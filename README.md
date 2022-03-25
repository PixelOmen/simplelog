# Simplelog

A single package that provides a simple, easy to use logger that can toggle between different log levels (DEBUG, INFO, WARNING, ERROR, FATAL). Most of the actual logging is done with the stdlib log.Logger, simplelog mainly controls which messages get written to logs via log levels.

# Install
## Clone Repo
```
git clone https://github.com/pixelomen/simplelog
```
## Go Get
```
go get github.com/pixelomen/simplelog
```

# Usage
```go
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
```
## testlog.txt output
```
2022/01/14 12:13:03 DEBUG: testapp.go:13: First msg
2022/01/14 12:13:03 INFO: testapp.go:14: First msg
2022/01/14 12:13:03 WARNING: testapp.go:15: First msg
2022/01/14 12:13:03 ERROR: testapp.go:16: First msg
2022/01/14 12:13:03 INFO: testapp.go:19: Another pointer
2022/01/14 12:13:03 INFO: testapp.go:20: --------------
2022/01/14 12:13:03 WARNING: testapp.go:25: Warnings and higher
2022/01/14 12:13:03 ERROR: testapp.go:26: Warnings and higher
```
## stdout output
```
INFO: Just a prefix
```
---
##
## Contributing
Pull requests are welcome. 

## License
[MIT](https://choosealicense.com/licenses/mit/)