package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
)

const (
	fatal     = 5
	errorMsg  = 4
	warn      = 3
	important = 2
	info      = 1
)

var logColors = map[int]*color.Color{
	fatal:     color.New(color.FgRed).Add(color.Bold),
	errorMsg:  color.New(color.FgRed),
	warn:      color.New(color.FgYellow),
	important: color.New(color.Bold),
}

type logger struct {
	sync.Mutex

	verbose bool
}

func (l *logger) SetVerbose(f bool) {
	l.verbose = f
}

func (l *logger) Log(level int, format string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()
	if level < errorMsg && !l.verbose {
		return
	}

	if c, ok := logColors[level]; ok {
		c.Printf(format, args...)
	} else {
		fmt.Printf(format, args...)
	}

	if level == fatal {
		os.Exit(1)
	}
}

func (l *logger) Fatal(format string, args ...interface{}) {
	l.Log(fatal, format, args...)
}

func (l *logger) Error(format string, args ...interface{}) {
	l.Log(errorMsg, format, args...)
}

func (l *logger) Warn(format string, args ...interface{}) {
	l.Log(warn, format, args...)
}

func (l *logger) Important(format string, args ...interface{}) {
	l.Log(important, format, args...)
}

func (l *logger) Info(format string, args ...interface{}) {
	l.Log(info, format, args...)
}
