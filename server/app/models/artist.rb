class Artist < ActiveRecord::Base
  has_many :songs
  has_many :albums, ->{ uniq }, through: :songs
  has_many :links, class_name: "ArtistLink", foreign_key: "ref_id"
end
