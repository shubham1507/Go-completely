package main

import "fmt"

type Loglevel int

const (
	LevelTrace Loglevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var levelNames = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}

func (l Loglevel) String() string {
	if l < LevelTrace || l > LevelFatal {
		return "UNKNOWN"
	}
	return levelNames[l]
}

func printLog(level Loglevel) {
	fmt.Printf("Log level: %d %s\n", level, level.String())
}

func main() {
	printLog(LevelTrace)
	printLog(LevelDebug)
	printLog(LevelInfo)
	printLog(LevelWarn)
	printLog(LevelError)
	printLog(LevelFatal)
	printLog(11)

}
