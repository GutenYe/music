class CreateLinks < ActiveRecord::Migration
  def change
    create_table :links do |t|
      t.string :type
      t.string :name
      t.string :url
      t.references :ref

      t.timestamps
    end
    
    add_index :links, :type
  end
end
