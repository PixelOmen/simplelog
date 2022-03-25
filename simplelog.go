//Package simplelog provides a simple, easy to use logger that can toggle between
//different log levels (DEBUG, INFO, WARNING, ERROR, FATAL). Most of the actual
//logging is done with the stdlib log.Logger, simplelog mainly controls
//which messages get written to logs via log levels.
package simplelog

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"sync"
)

//Log level constants
const (
	DEBUG   = 0
	INFO    = 10
	WARNING = 20
	ERROR   = 30
	FATAL   = 40
)

//Pkg level lock for reading and writing to the package level and
//a map of all logs. allLogs tracks all existing Loggers.
var (
	pkgLock sync.Mutex
	allLogs = make(map[string]*Logger)
)

//getfileinfo gets the calling filename and line number.
func getfileinfo() string {
	_, filename, line, ok := runtime.Caller(2)
	if !ok {
		filename = "Unknown"
		line = 0
	}
	return fmt.Sprintf("%s:%d: ", filepath.Base(filename), line)
}

type Logger struct {
	name          string      //The same string used when calling Get()
	level         int         //Determines which of the loggers are allowed to write data
	debugLogger   *log.Logger //Independent loggers for Independent prefixes/loglevels
	infoLogger    *log.Logger //Independent loggers for Independent prefixes/loglevels
	warningLogger *log.Logger //Independent loggers for Independent prefixes/loglevels
	errorLogger   *log.Logger //Independent loggers for Independent prefixes/loglevels
	fatalLogger   *log.Logger //Independent loggers for Independent prefixes/loglevels
	logfileinfo   bool        //Whether or not to log the calling file and line number
	lock          sync.Mutex  //Extra concurrency protection
}

//SetLevel sets the log level for the respective Logger.
//Each level is an int that can be set via pkg constants
//e.g. (simplelog.DEBUG) or via literal ints. Messages sent
//to Loggers with a log level lower than the current level are
//not written.
func (logger *Logger) SetLevel(level int) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	logger.level = level
}

//Log level = 0 | simplelog.DEBUG
func (logger *Logger) Debug(msg string) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	if logger.level > DEBUG {
		return
	}
	if logger.logfileinfo {
		msg = getfileinfo() + msg
	}
	logger.debugLogger.Println(msg)
}

//Log level = 10 | simplelog.INFO
func (logger *Logger) Info(msg string) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	if logger.level > INFO {
		return
	}
	if logger.logfileinfo {
		msg = getfileinfo() + msg
	}
	logger.infoLogger.Println(msg)
}

//Log level = 20 | simplelog.WARNING
func (logger *Logger) Warning(msg string) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	if logger.level > WARNING {
		return
	}
	if logger.logfileinfo {
		msg = getfileinfo() + msg
	}
	logger.warningLogger.Println(msg)
}

//Log level = 30 | simplelog.ERROR
func (logger *Logger) Err(msg string) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	if logger.level > ERROR {
		return
	}
	if logger.logfileinfo {
		msg = getfileinfo() + msg
	}
	logger.errorLogger.Println(msg)
}

func (logger *Logger) Fatal(msg string) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	if logger.level > FATAL {
		return
	}
	if logger.logfileinfo {
		msg = getfileinfo() + msg
	}
	logger.fatalLogger.Fatal(msg)
}

/*
New is the constructor for simplelog. It returns a pointer to a lvlLogger.
	name - Can be used to get a new pointer to an existing log via simplelog.Get(name)
	dest - Sets the destination to which log messages will be written
	logfileinfo - Whether or not to include filename:line in message (e.g. main.go:30)
	flags - The same flags you would pass into the stdlib log.New()
		Defaults to "log.Ldate | log.Ltime | log.Lmsgprefix" if nothing is passed.
		`log.Lshortfile` will always report as this pkg, so use logfileinfo the param instead.

*/
func New(name string, dest io.Writer, logfileinfo bool, flags ...int) *Logger {
	pkgLock.Lock()
	defer pkgLock.Unlock()
	_, alreadyExists := allLogs[name]
	if alreadyExists {
		panic(fmt.Sprintf("simplelog: Unable to create logger with name \"%s\", name already in use", name))
	}
	var logflags int
	if flags == nil {
		logflags = log.Ldate | log.Ltime | log.Lmsgprefix
	} else {
		logflags = flags[0]
	}
	debugLogger := log.New(dest, "DEBUG: ", logflags)
	infoLogger := log.New(dest, "INFO: ", logflags)
	warningLogger := log.New(dest, "WARNING: ", logflags)
	errorLogger := log.New(dest, "ERROR: ", logflags)
	fatalLogger := log.New(dest, "FATAL: ", logflags)
	returnLogger := &Logger{
		name:          name,
		debugLogger:   debugLogger,
		infoLogger:    infoLogger,
		warningLogger: warningLogger,
		errorLogger:   errorLogger,
		fatalLogger:   fatalLogger,
		logfileinfo:   logfileinfo,
	}
	allLogs[name] = returnLogger
	return returnLogger
}

//Get returns a reference to an existing Logger if one exists, otherwise nil.
//Uses the name string that was used to create the log via simplelog.New()
func Get(name string) *Logger {
	pkgLock.Lock()
	defer pkgLock.Unlock()
	if foundlog, isfound := allLogs[name]; isfound {
		return foundlog
	} else {
		return nil
	}
}
