package audio

import (
	"fmt"
)

const (
	MESSAGE_UNKOWN = 0
	MESSAGE_EOS    = 1 << iota
	MESSAGE_ERROR
)

const (
	STATE_STOP     = iota
	STATE_PLAYING
	STATE_PAUSED
)

type Backend interface {
	New(string) Audio
}

type Audio interface {
	GetName() string
	Play(uri string)
	Pause()
	Resume()
	Stop()
	GetState() int
	OnMessage(cb_func, param0 interface{})
}

var backends = make(map[string]Backend)

func Register(name string, backend Backend) {
	if backend == nil {
		panic("audio: Register backend is nil")
	}
	if _, dup := backends[name]; dup {
		panic("audio: Register called twice for backend " + name)
	}
	backends[name] = backend
}

func NewAudio(backendName string) Audio {
	backend, ok := backends[backendName]
	if !ok {
		fmt.Printf("audio: unknown backend %q (forgotten import?)\n", backendName)
	}
	return backend.New(backendName)
}
