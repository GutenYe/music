class CreateSongs < ActiveRecord::Migration
  def change
    create_table :songs do |t|
      t.references :album
      t.references :artist
      t.string :mbid
      t.string :title
      t.integer :tracknum
      t.float :length
      t.string :file
      t.integer :played_count, default: 0

      t.timestamps
    end
    add_index :songs, :album_id
  end
end
