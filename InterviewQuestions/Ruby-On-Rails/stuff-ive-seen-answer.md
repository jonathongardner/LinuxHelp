```ruby
class User < ApplicationRecord
  ...
  def total_likes
    likes = Post.where(owner: self.username).joins(:post_metadata).sum(:likes_count)
    {
      likes: likes,
      mega_likes: likes > 100_000
    }
  end
  # OR
  def total_likes
    likes = Post.where(owner: self.username).includes(:post_metadata).sum do |post|
      post.post_metadata&.likes_count || 0
    end
    {
      likes: likes,
      mega_likes: likes > 100_000
    }
  end
  ...
end
```
Big
  - N + 1 Queries
  - Passed username when instance method
  - Used if statement to set boolean

Small
  - Used `set` & `get` in method name
  - Didnt use SQL method
