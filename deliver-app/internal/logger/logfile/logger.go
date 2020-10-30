package logfile

import (
	"log"
	"os"
)

type logger struct {
	fileName string
}

type Logger interface {
	Init() *os.File
}

// NewLogger initialize logfile logger
func NewLogger(fileName string) Logger {
	return &logger{
		fileName: fileName,
	}
}

func (l *logger) Init() *os.File{
	f, err := os.OpenFile(l.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Can't create log file")
	}
	return f
}