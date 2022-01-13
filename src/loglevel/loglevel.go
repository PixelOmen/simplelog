package loglevel

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

func getfileinfo() (filenameLine string) {
	_, filename, line, ok := runtime.Caller(2)
	if !ok {
		filename = "Unknown"
		line = 0
	}
	return fmt.Sprintf("%s:%d: ", filepath.Base(filename), line)
}

type loglevel struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	logfileinfo   bool
	lock          sync.Mutex
	level         int
}

func (logger *loglevel) SetLevel(level int) {
	logger.level = level
}

func (logger *loglevel) Debug(msg string) {
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

func (logger *loglevel) Info(msg string) {
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

func (logger *loglevel) Warning(msg string) {
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

func (logger *loglevel) Err(msg string) {
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

func New(dest io.Writer, logfileinfo bool, flags ...int) loglevel {
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
	return loglevel{
		debugLogger:   debugLogger,
		infoLogger:    infoLogger,
		warningLogger: warningLogger,
		errorLogger:   errorLogger,
		logfileinfo:   logfileinfo,
	}
}
