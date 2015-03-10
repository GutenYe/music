package main

import (
	"./api/http_api"
	_ "./audio/gstreamer"
	. "./globals/ui"
	. "./globals/player"
	. "./globals/rc"
	"./rc"
	"./player"
	"github.com/ziutek/glib"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"github.com/ogier/pflag"
	"github.com/GutenYe/tagen.go/os2"
)
import . "github.com/GutenYe/tagen.go/pd"; var _ = Pd

const VERSION = "1.0.0"
var HOME = "home"
var	homeDir = filepath.Join(HOME, ".milk")
var	stateFile = filepath.Join(homeDir, "state")

var USAGE = `
$ milk-daemon [options]

OPTIONS
	-v, --version
	-h, --help
	--http=:3001       # http-api address
`
func initHomeDir() {
	if os2.IsNotExist(homeDir) {
		Ui.Printf("init home config directory %s", homeDir)
		e := os.MkdirAll(homeDir, 0755)
		if e != nil { Ui.Printf("initHomeDir error %s", e) }
	}
}

func trapSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for {
			<-c
			Ui.Printf("saving state to %s", stateFile)
			e := Player.SaveState(stateFile)
			if e != nil { Ui.Printf("Player.SaveState error: %s", e) }
			os.Exit(0)
		}
	}()
}

func loadState() {
	if os2.IsExist(stateFile) {
		Ui.Printf("loading state from %s", stateFile)
		var e error
		Player, e = player.LoadState(stateFile)
		if e != nil { Ui.Printf("player.LoadState error: %s", e) }
	} else {
		Player = player.NewPlayer("gstreamer")
	}
}

func main() {
	pflag.Usage = func() {
		Ui.Print(USAGE)
	}
	var version = pflag.BoolP("version", "", false, "Print version number")
	var http = pflag.StringP("http", "", ":3002", "http-api address")
	var server = pflag.StringP("server", "", "http://localhost:3000", "server url")
	var logfile = pflag.StringP("log", "", "", "log file")
	pflag.Parse()

	if *version {
		Ui.Println(VERSION)
		return
	}

	if *logfile != "" {
		w, e := os.OpenFile(*logfile, os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_TRUNC, 0644)
		if e != nil { panic(e) }
		defer w.Close()
		Ui = log.New(w, "", log.Ldate | log.Ltime)
	} else {
		Ui = log.New(os.Stdout, "", 0)
	}

	initHomeDir()
	loadState()
	Rc = &rc.Rc{SERVER_URL: *server}
	trapSignals()

	go http_api.Start(*http)
	glib.NewMainLoop(nil).Run()
}
