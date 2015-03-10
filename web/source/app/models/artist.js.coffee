A.Artist = DS.Model.extend
  name: DS.attr("string")
  gender: DS.attr("string")
  
  songs: DS.hasMany("A.Song")
  albums: DS.hasMany("A.Album")
