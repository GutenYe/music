A.Song = DS.Model.extend
  title: DS.attr("string")
  tracknum: DS.attr("number")
  length: DS.attr("number")
  fileUrl: DS.attr("string")
  playedCount: DS.attr("string")

  artist: DS.belongsTo("A.Artist")
  album: DS.belongsTo("A.Album")
  links: DS.hasMany("A.Link")
