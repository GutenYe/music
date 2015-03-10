package http_api

import (
	"os"
	"strings"
	"net"
	"net/http"
	"../../api"
	. "../../globals/ui"
	. "../../globals/player"
	"encoding/json"
	"fmt"
)

type Message struct {
	Error bool
	Message string
	Current int
	Song  song.Song   `json:,omitempty`
	Songs []song.Song `json:,omitempty`
}

func renderJSON(w http.ResponseWriter, msg Message) {
	b, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(b)))
	fmt.Fprint(w, string(b))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		e := api.Play()
		var msg Message
		if e != nil {
			msg = Message{Error: true, Message: e}
		} else {
			msg = Message{Song: Player.CurrentSong()}
		}
		renderJSON(w, msg)
	}
}

func PlayOrResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		if Player.IsPaused() {
			Player.Resume()
		} else {
			api.Play()
		}
		renderJSON(w, Message{Song: Player.CurrentSong()})
	}
}

func PauseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		Player.Pause()
	}
}

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		Player.Resume()
	}
}

func StopHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		Player.Stop()
	}
}

func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderJSON(w, Message{Song: Player.CurrentSong()})
	}
}

func PlaylistHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderJSON(w, Message{Songs: Player.CurrentSongs(), Current: Player.CurrentPosition()})
	}
}

func NextHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		if r.FormValue("album") != "" { 
			Player.PlayNextAlbum()
		} else {
			Player.PlayNext()
		}
		renderJSON(w, Message{song: Player.CurrentSong()})
	}
}

func PrevHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		if r.FormValue("album") != "" { 
			Player.PlayPrevAlbum()
		} else {
			Player.PlayPrev()
		}
		renderJSON(w, Message{song: Player.CurrentSong()})
	}
}

func Start(addr string) {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/play", PlayHandler)
	http.HandleFunc("/play_or_resume", PlayOrResumeHandler)
	http.HandleFunc("/pause", PauseHandler)
	http.HandleFunc("/resume", ResumeHandler)
	http.HandleFunc("/stop", StopHandler)
	http.HandleFunc("/current", CurrentHandler)
	http.HandleFunc("/next", NextHandler)
	http.HandleFunc("/prev", PrevHandler)
	http.HandleFunc("/playlist", PlaylistHandler)

	Ui.Printf("> HTTP API start at %s\n", addr) 

	if strings.HasPrefix(addr, "unix://") {
		addr = addr[7:]
		os.Remove(addr)
		l, err := net.Listen("unix", addr)
		if err != nil { panic(err) }
		os.Chmod(addr, 0777)
		Ui.Fatal(http.Serve(l, http.DefaultServeMux))
	} else {
		Ui.Fatal(http.ListenAndServe(addr, nil))
	}
}
