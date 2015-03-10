package song

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	. "../globals/rc"
	//"fmt"
)

type Song struct {
	Id      int
	Title   string
	Album   string
	Artist  string
	FileUrl string `json:"file_url"`
}

func All() (songs []Song, err error) {
	resp, err := http.Get(Rc.SERVER_URL+"/songs2")
	if err != nil { return nil, err }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { return nil, err }

	err = json.Unmarshal(body, &songs)
	if err != nil { return nil, err }
	return songs, nil
}
