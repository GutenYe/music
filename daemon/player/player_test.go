package player

import (
	"testing"
	"../song"
	_ "../audio/dummy"
)

var songs = []song.Song{
	song.Song{Id: 1, Title: "a", Album: "b", Artist: "c", FileUrl: "d"},
	song.Song{Id: 2, Title: "a2", Album: "b", Artist: "c2", FileUrl: "d2"},
	song.Song{Id: 3, Title: "a3", Album: "b3", Artist: "c3", FileUrl: "d3"},
}

func TestAddSongs(t *testing.T) {
	p := NewPlayer("dummy")
	p.AddSongs(songs...)

	in, out := p.current, p.queue.Front()
	if in != out {
		t.Fatalf("AddSongs: current is %p, want %p\n", in, out)
	} 

	in2, out2 := p.queue.Len(), 3
	if in2 != out2 {
		t.Fatalf("AddSongs: queue's length is %d, want %d\n", in2, out2)
	}
}

func TestClearSongs(t *testing.T) {
	p := NewPlayer("dummy")
	p.AddSongs(songs...)
	p.ClearSongs()

	in := p.current
	if in != nil {
		t.Fatalf("ClearSongs: current is %q, want %q\n", in, nil)
	}

	in2, out2 := p.queue.Len(), 0
	if in2 != out2 {
		t.Fatalf("ClearSongs: queue.Len() is %d, want %d\n", in2, out2)
	}
}

func TestGetNext(t *testing.T) {
	p := NewPlayer("dummy")
	p.AddSongs(songs...)

	// current is nil
	p.current = nil
	in := p.GetNext()
	if in != nil {
		t.Errorf("GetNext: current is nil => %q, want %q", in, nil)
	}

	// current at first song
	p.current = p.queue.Front()
	in2, out2 := p.GetNext(), p.queue.Front().Next()
	if in2 != out2 {
		t.Errorf("GetNext: current at first song => %q, want %q", in2, out2)
	}

	// current at last song
	p.current = p.queue.Back()
	in3, out3 := p.GetNext(), p.queue.Front()
	if in3 != out3 {
		t.Errorf("GetNext: current at last song => %q, want %q", in3, out3)
	}
}

