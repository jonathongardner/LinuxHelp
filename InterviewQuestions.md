## Question 1
You have:
```ruby
class Post < ApplicationRecord
  has_one :post_metadata
  ...
end
```
```ruby
class PostMetadata < ApplicationRecord
  belongs_to :post
  ...
end
```
```ruby
class User < ApplicationRecord
  ...
  def get_total_likes(username)
    likes = 0
    mega_likes = false
    Post.where(owner: username).each do |p|
      if p.post_metadata&.likes_count
        likes += p.post_metadata.likes_count
      end
    end
    if likes > 100_000
      mega_likes = true
    end
    {
      likes: likes,
      mega_likes: mega_likes
    }
  end
  ...
end
```
```ruby
class SomeController < ApplicationController
  ...
  def set_like_information
    @like_information = current_user.get_total_likes(current_user.username)
  end
  ...
end
```
Whats wrong with this?

## Question 2
```ruby
def special_method1(array)
  to_return = array.reduce('') do |acc, value|
    next acc if value == ' '
    "#{acc}#{value}"
  end
  "special_method1: #{to_return}"
end

def special_method2(array)
  to_return = array.reduce('') do |acc, value|
    break acc if value == ' '
    "#{acc}#{value}"
  end
  "special_method2: #{to_return}"
end

def special_method3(array)
  to_return = array.reduce('') do |acc, value|
    return acc if value == ' '
    "#{acc}#{value}"
  end
  "special_method3: #{to_return}"
end
```
What is returned when the following are called
```ruby
# ['t', 'h', 'i', 's', ' ', 'i', 's', ' ', 'a', ' ', 't', 'e', 's', 't']
puts special_method1('this is a test'.split(''))
puts special_method2('this is a test'.split(''))
puts special_method3('this is a test'.split(''))

puts special_method1('this_is_a_test'.split(''))
puts special_method2('this_is_a_test'.split(''))
puts special_method3('this_is_a_test'.split(''))
```

## Question 3
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

## Question 4
```ruby
def some_method(*options1, **options2)
  puts options1.to_s
  puts options2.to_s
end
```
Whats in option1 and option2?
