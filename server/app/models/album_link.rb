class AlbumLink < Link
  belongs_to :album, foreign_key: "ref_id"
end
