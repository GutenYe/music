class SongLink < Link
  belongs_to :song, foreign_key: "ref_id"
end
