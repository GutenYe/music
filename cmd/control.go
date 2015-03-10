package main

import (
	"net/http"
	"io/ioutil"
	. "./globals/ui"
	. "./globals/rc"
	"encoding/json"
)

type Song struct {
	Id      int
	Title   string
	Album   string
	Artist  string
	FileUrl string `json:"file_url"`
}

type Message struct {
	Error bool
	Message string
	Current int
	Song  song.Song   `json:,omitempty`
	Songs []song.Song `json:,omitempty`
}

func parseMessage(resp *http.Response) (msg Message) {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &msg)
	return msg
}

func Put(url string) (resp *http.Response, err error)  {
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, nil)
	resp, err = client.Do(req)
	return resp, err
}

func Play() {
	r, _ := Put(Rc.DAEMON_URL+"/play_or_resume")
	s := parseMessage(r).Song
	Ui.Printf("playing [%s - %s] %s\n", s.Artist, s.Album, s.Title)
}

func Pause() {
	Put(Rc.DAEMON_URL+"/pause")
}

func Stop() {
	Put(Rc.DAEMON_URL+"/stop")
}

func Current() {
	r, _ := http.Get(Rc.DAEMON_URL+"/current")
	s := parseMessage(r).Song
	Ui.Printf("[%s - %s] %s\n", s.Artist, s.Album, s.Title)
}

func Next(album string) {
	url := Rc.DAEMON_URL+"/next"
	if album != "" { 
		url += "?album=true"
	}
	r, _ := Put(url)
	s := parseMessage(r).Song
	Ui.Printf("playing [%s - %s] %s\n", s.Artist, s.Album, s.Title)
}

func Prev(album string) {
	url := Rc.DAEMON_URL+"/prev"
	if album != "" { 
		url += "?album=true"
	}
	r, _ := Put(url)
	s := parseMessage(r).Song
	Ui.Printf("playing [%s - %s] %s\n", s.Artist, s.Album, s.Title)
}

func Playlist() {
	r, _ := http.Get(Rc.DAEMON_URL+"/current")
	msg := parseMessage(r)
	var prefix string
	for i, s := range msg.Songs {
		if i == msg.Current {
			prefix = "**"
		} else { 
			prefix = "#"
		}
		Ui.Printf("%s%2d [%s - %s] %s\n", prefix, i+1, s.Artist, s.Album, s.Title)
	}
}
