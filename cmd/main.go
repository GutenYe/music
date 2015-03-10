package main

import (
	. "./globals/rc"
	. "./globals/ui"
	"github.com/GutenYe/tagen.go/os2"
	"github.com/ogier/pflag"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
	"path/filepath"
)
import . "github.com/GutenYe/tagen.go/pd"; var _ = Pd

const VERSION = "1.0.0"
var HOME = "home"
var homeRc = filepath.Join(HOME, ".milkrc")

var USAGE = `
$ milk <cmd> [options]

COMMAND:
	upload <dir/file ...>

	play
	pause
	stop
	next
	prev
	current

GENERIC OPTIONS
	-v, --version
	-h, --help
`
func main() {
	pflag.Usage = func() {
		Ui.Print(USAGE)
	}
	var version = pflag.BoolP("version", "v", false, "Print version number")
	var server = pflag.StringP("server", "", "", "server url")
	var daemon = pflag.StringP("daemon", "", "", "daemon url")
	pflag.Parse()
	if *version {
		Ui.Printf("milk %s", VERSION)
		return
	}

	Ui = log.New(os.Stdout, "", 0)
	if os2.IsExist(homeRc) {
		d, e := ioutil.ReadFile(homeRc)
		if e != nil {
			Ui.Printf("can't read homerc from %s", homeRc)
			return
		}
		goyaml.Unmarshal(d, &Rc)
	}
	if *server != "" { Rc.SERVER_URL = *server }
	if *daemon != "" { Rc.DAEMON_URL = *daemon }

	switch pflag.Arg(0) {
	case "upload":
		Upload(pflag.Args()[1:]...)
	case "play":
		Play()
	case "pause":
		Pause()
	case "stop":
		Stop()
	case "current":
		Current()
	case "next":
		Next(pflag.Arg(1))
	case "prev":
		Prev(pflag.Arg(1))
	case "playlist":
		Playlist()
	default:
		pflag.Usage()
	}
}
