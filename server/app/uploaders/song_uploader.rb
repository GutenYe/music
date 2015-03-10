class SongUploader < CarrierWave::Uploader::Base
  include CarrierWave::Video

  storage :file

  def store_dir
    "uploads/music/#{model.id}"
  end

  def cache_dir
    "/tmp/carrier_wave"
  end

  version :mp3 do
    process :encode_video => [:mp3]
    def full_filename(original)
      "#{File.basename(original, File.extname(original))}.mp3"
    end
  end

  def extension_white_list
    %w(flac mp3 ogg wav)
  end
end
