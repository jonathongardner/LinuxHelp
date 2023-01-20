class Init < ActiveRecord::Migration[5.2]
  def change
    create_table :posts do |t|
      t.string :message
      t.string :for
      t.string :owner, index: true
    end
    create_table :post_metadatas do |t|
      t.bigint :likes_count
      t.belongs_to :post
    end
    create_table :users do |t|
      t.string :username
    end
  end
end
