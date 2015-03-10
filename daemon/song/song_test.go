package song

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	. "../globals/rc"
	"../rc"
)

var songs_bytes = `[
	{"id":1,"title":"a","album":"b","artist":"c","file_url":"d"},
	{"id":2,"title":"a1","album":"b1","artist":"c1","file_url":"d1"}]`
var songs_tests = []Song{
	Song{Id: 1, Title: "a", Album: "b", Artist: "c", FileUrl: "d"},
	Song{Id: 2, Title: "a1", Album: "b1", Artist: "c1", FileUrl: "d1"},
}

func Songs2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, songs_bytes)
}

func startServer() {
	Rc = &rc.Rc{}

	mux := http.NewServeMux()
	mux.HandleFunc("/songs2", Songs2Handler)
	ts := httptest.NewServer(mux)
	Rc.SERVER_URL = ts.URL
}

func TestAll(t *testing.T) {
	startServer()
	songs, e := All()

	if e != nil {
		t.Fatal(e)
	}

	for i, v := range songs {
		out := songs_tests[i]
		if v != out {
			t.Fatalf("%d. All() => %q, want %q", i, v, out)
		}
	}
}
