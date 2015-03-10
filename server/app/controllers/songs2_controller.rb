class Songs2Controller < ApplicationController
  def index
    @songs = Song.includes(:artist, :album)

    render json: @songs,
      :root => false,
     	:each_serializer => Song2Serializer
  end
end
