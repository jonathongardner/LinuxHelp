## Question 1
```ruby
my_hash = { key1: 'value1', 'key2' => 'value2' }
```
What would `my_hash['key1']` return?






## Question 2
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




## Question 3
```ruby
def special_method(array)
  to_return = array.map do |value|
    next 47 if value == 3
    break 47 if value == 4
    return 47 if value == 5
    value
  end
  "special_method: #{to_return}"
end
```
What is returned when the following are called
```ruby
special_method([2, 3, 7])
special_method([2, 4, 7])
special_method([2, 5, 7])
```





## Question 4
```ruby
result1 = ['a', 'b', 'c', 'd'].each_with_object({}) do |value, acc|
  acc.merge(value => value)
end

result2 = ['a', 'b', 'c', 'd'].reduce({}) do |acc, value|
  acc.merge(value => value)
end

puts result1
puts result2
```
Whats in result1 and result2?





## Question 5
```ruby
def some_method(*options1, **options2)
  puts options1.to_s
  puts options2.to_s
end
```
```ruby
some_method('a', 'b', c: 'c', 'd' => 'd', e: 'e', f: 'f')
```
Whats in option1 and option2?
