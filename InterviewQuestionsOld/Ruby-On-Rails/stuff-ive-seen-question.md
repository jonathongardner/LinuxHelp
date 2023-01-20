You have:
```ruby
class Post < ApplicationRecord
  has_one :post_metadata
  #...
end
```
```ruby
class PostMetadata < ApplicationRecord
  belongs_to :post
  #...
end
```
```ruby
class User < ApplicationRecord
  #...
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
  #...
end
```
```ruby
class SomeController < ApplicationController
  #...
  def set_like_information
    @like_information = current_user.get_total_likes(current_user.username)
  end
  #...
end
```
Whats wrong with this/How would you refactor this?
