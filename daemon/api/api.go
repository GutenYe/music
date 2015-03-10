package api

import (
	"../song"
	. "../globals/player"
)


func Play() error {
	songs, e := song.All()
	Player.ClearSongs()
	Player.AddSongs(songs...)
	Player.Play()
	return e
}
