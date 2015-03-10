class ArtistLink < Link
  belongs_to :artist, foreign_key: "ref_id"
end
