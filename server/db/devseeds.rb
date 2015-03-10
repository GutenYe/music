[ # Artist
  ["a-b", "Lenka"],
  ["a-c", "Guten"],
].each.with_index(1) {|(mbid, name),i|
  r = Artist.new(mbid: mbid, name: name)
  #r.id = i
  r.save!
}

[ # ArtistLink
  ["lastfm", "http://artist/lenka", "Lenka"],
  ["lastfm", "http://artist/guten", "Guten"],
].each.with_index(1) {|(name, url, artist),i|
  r = ArtistLink.new(name: name, url: url)
  r.artist = Artist.find_by_name(artist)
  #r.id = i
  r.save!
}

[ # Album
  ["b-a", "Trouble", "2011-10-1", "Lenka"],
  ["b-b", "Life", "2013-1-1", "Guten"],
].each.with_index(1) {|(mbid, title, date, artist),i|
  r = Album.new(mbid: mbid, title: title, date: date)
  #r.id = i
  r.save!
}

[ # AlbumLink
  ["lastfm", "http://album/trouble", "Trouble"],
  ["lastfm", "http://album/life", "Life"],
].each.with_index(1) {|(name, url, album),i|
  r = AlbumLink.new(name: name, url: url)
  r.album = Album.find_by_title(album)
  #r.id = i
  r.save!
}

[ # Song
  ["c-a", 1, "Trouble Is A Friend", 121, "Lenka", "Trouble", "a.flac"],
  ["c-b", 2, "The Show", 182,            "Lenka", "Trouble", "b.flac"],
  ["c-c", 3, "Knock Knock", 199,         "Lenka", "Trouble", "c.flac"],
  ["c-d", 1, "Life is Hard", 221,        "Guten", "Life", "d.flac"],
].each.with_index(1) {|(mbid, tracknum, title, length, artist, album, file),i|
  r = Song.new(mbid: mbid, tracknum: tracknum, title: title, length: length)
  r.artist = Artist.find_by_name(artist)
  r.album = Album.find_by_title(album)
  r[:file] = file
  #r.id = i
  r.save!
}

[ # SongLink
  ["lastfm", "http://song/trouble-is-a-friend", "Trouble Is A Friend"],
  ["lastfm", "http://song/the-show", "The Show"],
].each.with_index(1) {|(name, url, song),i|
  r = SongLink.new(name: name, url: url)
  r.song = Song.find_by_title(song)
  #r.id = i
  r.save!
}
