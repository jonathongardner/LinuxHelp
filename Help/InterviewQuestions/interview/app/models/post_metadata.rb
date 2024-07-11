class PostMetadata < ApplicationRecord
  belongs_to :post
  # fields:
  # likes: bigint
  # post_id: bigint fk to post.id
  # ...
end
