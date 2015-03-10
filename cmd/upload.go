package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"io/ioutil"
	. "./globals/ui"
	. "./globals/rc"
	"os"
	//"fmt"
)

var SONG_EXTS = []string{"flac"}

func Upload(paths ...string) {

	for _, path := range paths {
		filepath.Walk(path, func(p string, i os.FileInfo, err error) error {
			if i.IsDir() { return nil }

			isIncluded := false
			for _, v := range SONG_EXTS {
				if filepath.Ext(p) == "."+v  {
					isIncluded = true
					break
				}
			}
			if !isIncluded { return nil }

			doUpload(p)

			return nil
		})
	}
}

func doUpload(file string) {
	Ui.Printf("Uploading %s", file)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	part, _ := w.CreateFormFile("song[file]", filepath.Base(file))
	data, err := ioutil.ReadFile(file)
	if err != nil { panic(err) }
	part.Write(data)
	w.Close()

	http.Post(Rc.SERVER_URL+"/songs", w.FormDataContentType(), &b)
}
