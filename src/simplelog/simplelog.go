package simplelog

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	DEBUG   = 0
	INFO    = 10
	WARNING = 20
	ERROR   = 30
)

var (
	pkgLock sync.Mutex
	allLogs = make(map[string]*lvlLog)
)

func getfileinfo() (filenameLine string) {
	_, filename, line, ok := runtime.Caller(2)
	if !ok {
		filename = "Unknown"
		line = 0
	}
	return fmt.Sprintf("%s:%d: ", filepath.Base(filename), line)
}

type lvlLog struct {
	name          string
	level         int
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	logfileinfo   bool
	lock          sync.Mutex
}

func (logger *lvlLog) SetLevel(level int) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	logger.level = level
}

func (logger *lvlLog) Debug(msg string) {
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

func (logger *lvlLog) Info(msg string) {
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

func (logger *lvlLog) Warning(msg string) {
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

func (logger *lvlLog) Err(msg string) {
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

func New(name string, dest io.Writer, logfileinfo bool, flags ...int) *lvlLog {
	pkgLock.Lock()
	defer pkgLock.Unlock()
	_, alreadyExists := allLogs[name]
	if alreadyExists {
		panic(fmt.Sprintf("levelslog: Unable to create logger with name \"%s\", name already in use", name))
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
	errorLogger := log.New(dest, "Error: ", logflags)
	returnlogger := &lvlLog{
		name:          name,
		debugLogger:   debugLogger,
		infoLogger:    infoLogger,
		warningLogger: warningLogger,
		errorLogger:   errorLogger,
		logfileinfo:   logfileinfo,
	}
	allLogs[name] = returnlogger
	return returnlogger
}

func Get(name string) *lvlLog {
	pkgLock.Lock()
	defer pkgLock.Unlock()
	if foundlog, isfound := allLogs[name]; isfound {
		return foundlog
	} else {
		return nil
	}
}
