package player

import (
	"../audio"
	"../song"
	"container/list"
	"fmt"
	. "../globals/ui"
	"encoding/json"
	"io/ioutil"
)

type Player struct {
	audio    audio.Audio     // the backend of the player
	queue    *list.List
	current  *list.Element   // current position in the queue. init by AddSongs
	repeat   bool
	single   bool
	random   bool
}

type State struct {
	Queue []song.Song
	Current int
	Audio string
	Repeat bool
	Single bool
	Random bool
}

func NewPlayer(backendName string) *Player {
	p := &Player{queue: list.New()}
	if backendName != "" { 
		p.SetAudio(backendName)
	}

	return p
}

func LoadState(file string) (*Player, error) {
	b, e := ioutil.ReadFile(file)
	if e != nil { return nil, e  }

	p := NewPlayer("")
	e = json.Unmarshal(b, &p)
	if e != nil { return nil, e }

	return p, nil
}

func (p *Player) SaveState(file string) error {
	b, e := json.Marshal(p)
	if e != nil { return e }
	return ioutil.WriteFile(file, b, 0644)
}

func (p *Player) SetAudio(backendName string) {
	p.audio = audio.NewAudio(backendName)
	p.audio.OnMessage((*Player).OnMessage, p)
}

/* event  */

func (p *Player) OnMessage(msg int) {
	switch msg {
	case audio.MESSAGE_EOS:
		// single && repeat
		if p.single && p.repeat {
			p.Play()
		// single 
		} else if p.single {
			p.Stop()
		// !single - next
		} else {
			// random
			// TODO

			// end of queue and !repeat
			if p.current.Next() == nil && !p.repeat {
				p.Stop()
			} else {
				p.PlayNext()
			}
		}
	case audio.MESSAGE_ERROR:
		fmt.Println("audio ERROR")
	}
}

/* play & queue control */

func (p *Player) Play() {
	if p.current == nil {
		return
	}

	song := p.CurrentSong()
	Ui.Printf("playing %s\n", song.FileUrl)
	p.audio.Play(song.FileUrl)
}

func (p *Player) PlayNext() {
	if p.current == nil {
		return
	}

	p.current = p.GetNext()
	p.Play()
}

func (p *Player) PlayPrev() {
	if p.current == nil {
		return
	}

	p.current = p.GetPrev()
	p.Play()
}

func (p *Player) PlayNextAlbum() {
	if p.current == nil {
		return
	}

	p.current = p.GetNextAlbum()
	p.Play()
}

func (p *Player) PlayPrevAlbum() {
	if p.current == nil {
		return
	}

	p.current = p.GetPrevAlbum()
	p.Play()
}

func (p *Player) Pause() {
	p.audio.Pause()
}

func (p *Player) Resume() {
	p.audio.Resume()
}

func (p *Player) Stop() {
	p.audio.Stop()
}

/* queue control */
func (p *Player) AddSongs(songs ...song.Song) {
	for _, song := range songs {
		p.queue.PushBack(song)
	}

	if p.current == nil {
		p.current = p.queue.Front()
	}
}

func (p *Player) ClearSongs() {
	p.queue.Init()
	p.current = nil
}

func (p *Player) GetNext() *list.Element {
	if p.current == nil {
		return nil
	} 

	next := p.current.Next() 
	/* end of the queue */
	if next == nil {
		next = p.queue.Front()
	}

	return next
}

func (p *Player) GetPrev() *list.Element {
	if p.current == nil {
		return nil
	} 

	prev := p.current.Prev() 
	/* front of the queue */
	if prev == nil {
		prev = p.queue.Back()
	}

	return prev
}

func (p *Player) GetNextAlbum() *list.Element {
	if p.current == nil {
		return nil
	} 

	currentAlbum := p.CurrentSong().Album
	next := p.current.Next()
	for {
		/* looped once */
		if next == p.current {
			return nil
		}

		/* end of the queue */
		if next == nil {
			next = p.queue.Front()
		}

		/* find one */
		if next.Value.(song.Song).Album != currentAlbum {
			break
		}
	}

	return next
}

func (p *Player) GetPrevAlbum() *list.Element {
	if p.current == nil {
		return nil
	} 

	currentAlbum := p.CurrentSong().Album
	prev := p.current.Prev()
	for {
		/* looped once */
		if prev == p.current {
			return nil
		}

		/* front of the queue */
		if prev == nil {
			prev = p.queue.Back()
		}

		/* find one */
		if prev.Value.(song.Song).Album != currentAlbum {
			break
		}
	}

	return prev
}

// handle p.current == nil by your self.
func (p *Player) CurrentSong() song.Song {
	return p.current.Value.(song.Song)
}

// -1
func (p *Player) CurrentPosition() int {
	if p.current == nil {
		return -1
	}

	for v, i := p.queue.Front(), 0; v != nil; v, i = v.Next(), i + 1 {
		if v == p.current {
			return i
		}
	}

	return -1
}

func (p *Player) CurrentSongs() (songs []song.Song) {
	for v := p.queue.Front(); v != nil; v = v.Next() {
		songs = append(songs, v.Value.(song.Song))
	}

	return songs
}

/* other */
func (p *Player) IsStop() bool {
	return p.audio.GetState() == audio.STATE_STOP
}

func (p *Player) IsPlaying() bool {
	return p.audio.GetState() == audio.STATE_PLAYING
}

func (p *Player) IsPaused() bool {
	return p.audio.GetState() == audio.STATE_PAUSED
}

func (p *Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(State{
		Queue: p.CurrentSongs(),
		Current: p.CurrentPosition(),
		Audio: p.audio.GetName(),
		Repeat: p.repeat,
		Single: p.single,
		Random: p.random,
	})
}

func (p *Player) UnmarshalJSON(data []byte) error {
	s := State{}
	e := json.Unmarshal(data, &s)
	if e != nil { return e }

	p.SetAudio(s.Audio)
	p.repeat = s.Repeat
	p.single = s.Single
	p.random = s.Random
	for i, song := range s.Queue {
		v := p.queue.PushBack(song)
		if s.Current == i {
			p.current = v
		}
	}

	return nil
}
