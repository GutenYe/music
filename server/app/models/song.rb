class Song < ActiveRecord::Base
  belongs_to :album
  belongs_to :artist
  has_many :links, class_name: "SongLink", foreign_key: "ref_id"

  mount_uploader :file, SongUploader
end
