class Album < ActiveRecord::Base
  has_many :songs
  has_many :artists, -> { uniq }, through: :songs
  has_many :links, class_name: "AlbumLink", foreign_key: "ref_id"
end
