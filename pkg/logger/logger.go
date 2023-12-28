package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	LogFileName string
	Location    string
}

func LoadLogger(filename, location string) Logger {
	return Logger{LogFileName: filename, Location: location}
}

func (e *Logger) Log(s string, print bool) {
	f, err := os.OpenFile(e.Location+e.LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error Log() .. ")
		fmt.Println(err)
	}

	f.WriteString("[ " + time.Now().Format("01-02-2006 15:04:05") + " ] - " + s + "\n")

	if print {
		fmt.Println(s)
	}
}
