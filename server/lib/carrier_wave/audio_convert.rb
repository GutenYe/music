module CarrierWave
  module AudioConverter
    extend ActiveSupport::Concern

    included do
      begin
        require "streamio-ffmpeg"
      rescue LoadError => e
        require "voyeur"
      rescue LoadError => e
        e.message << " (You may need to install the streamio-ffmpeg or voyeur gem)"
        raise e
      end
    end

    module ClassMethods
      def encode_audio(format)
        process :encode_audio => format
      end
    end

    def encode_audio(format)
      # move upload to local cache
      cache_stored_file! if !cached?

      directory = File.dirname( current_path )

      # move upload to tmp file - encoding result will be saved to
      # original file name
      tmp_path   = File.join( directory, "tmpfile" )
      File.rename current_path, tmp_path

      # encode
      FFMPEG::Movie.new(tmp_path).transcode("movie.mp4")

      Voyeur::Video.new( filename: tmp_path ).convert( to: format.to_sym, output_filename: current_path )

      # because encoding video will change file extension, change it 
      # to old one
      fixed_name = File.basename(current_path, '.*') + "." + format.to_s
      File.rename File.join( directory, fixed_name ), current_path

      # delete tmp file
      File.delete tmp_path
    end

    private
      def prepare!
        cache_stored_file! if !cached?
      end
  end
end
