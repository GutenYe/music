package ui

import (
	"log"
	"os"
)

var std = log.New(os.Stdout, "", 0)

func Print(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
}
