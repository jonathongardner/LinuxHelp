class User < ApplicationRecord
  # fields:
  # username: string
  # ...

  # Gets total number of likes for all post by a given user
  def get_total_likes(username)
    likes = 0
    mega_likes = false
    Post.where(owner: username).each do |p|
      likes += p.post_metadata.likes_count
    end
    if likes > 100_000
      mega_likes = true
    end
    {
      likes: likes,
      mega_likes: mega_likes
    }
  end
end
