class CreateAlbums < ActiveRecord::Migration
  def change
    create_table :albums do |t|
      t.string :mbid
      t.string :title
      t.date :date

      t.timestamps
    end
  end
end
