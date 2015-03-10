A.Album = DS.Model.extend
  title: DS.attr("string")
  date: DS.attr("date")

  songs: DS.hasMany("A.Song")
  artists: DS.hasMany("A.Artist")
