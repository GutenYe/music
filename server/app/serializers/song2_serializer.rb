class Song2Serializer < ActiveModel::Serializer
  attributes :id, :title, :album, :artist, :file_url

  def album
    object.album.title
  end

  def artist
    object.artist.name
  end

  def file_url
    "#{$SERVER_ADDR}#{object.file.url}"
  end
end
