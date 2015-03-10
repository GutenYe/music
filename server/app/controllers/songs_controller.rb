class SongsController < ApplicationController
  def index
    @songs = if params["ids"]
      Song.where(id: params["ids"])
    else
      Song.all()
    end
    render json: @songs
  end

  def show
    @song = Song.find(params["id"])
    render json: @song
  end

  def create
    p = params[:song]
    file = p[:file].tempfile
    new_path = file.path+File.extname(p[:file].original_filename)
    File.symlink(file.path, new_path)

    AudioInfo.open(new_path) { |i|
      p[:title] ||= i.title  if i.title
      p[:tracknum] ||= i.tracknum if i.tracknum
      p[:length] ||= i.length if i.length
      p[:artist_id] ||= Artist.find_or_create_by_name(i.artist).id if i.artist
      p[:album_id]  ||= Album.find_or_create_by_title(i.album, date: i.date).id if i.album
    }

    @song = Song.new(p)
    if @song.save 
      puts "success"
    else
      puts "error"
    end
  end
end
