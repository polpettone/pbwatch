package internal

import (
	"io/ioutil"
	"log"
	"os"
)

type Logging struct {
	Stdout   *log.Logger
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	DebugLog *log.Logger
}

func openLogFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
    	log.Fatalf("error opening file: %v", err)
	}
	return f
}

func NewLogging(enabled bool) *Logging {
	var infoLog *log.Logger
	var debugLog *log.Logger

	stdout := log.New(os.Stdout, "", 0)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if enabled {
		infoLog = log.New(openLogFile("pbwatch_info.log"), "INFO\t", log.Ldate|log.Ltime)
		debugLog = log.New(openLogFile("pbwatch_debug.log"), "DEBUG\t", log.Ldate|log.Ltime)
	} else {
		infoLog = log.New(ioutil.Discard, "", 0)
		debugLog = log.New(ioutil.Discard, "", 0)
	}

	app := &Logging{
		Stdout:   stdout,
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		DebugLog: debugLog,
	}

	return app
}