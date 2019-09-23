## Question 1
```ruby
my_hash['key1'] # => nil
```
Whats the difference with string and symbol
How could you initialize this hash in RAILS so you could access it using string or symbol `HashWithIndifferentAccess`
Could I do that in ruby?


## Question 2
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

## Question 3
```ruby
special_method([2, 3, 7]) # => some_method: [2, 47, 7]
special_method([2, 4, 7]) # => some_method: 47
special_method([2, 5, 7]) # => 47
```
 - Next returns the value acc and goes to the next iterable value
 - Break returns the value acc and exits the loop
 - return returns the value and exits the method

## Question 4
```ruby
puts result1 # => {}
puts result2 # => {"a"=>"a", "b"=>"b", "c"=>"c", "d"=>"d"}
```
Fix
```ruby
result1 = ['a', 'b', 'c', 'd'].each_with_object({}) do |value, acc|
  acc.merge!(value => value)
end
```
 - `each_with_object` uses the value acc (so its like a reference)
 - `reduce` uses the value returned (in the case the last line but if next is used will use next value)

## Question 5
```ruby
options1 # => ["a", "b", {"d"=>"d"}]
options2 # => {:c=>"c", :e=>"e", :f=>"f"}
```
 - `**` only works with symbols (i.e. `{ **{ c: 'c', 'd' => 'd', e: 'e', f: 'f' } }` will return wrong argument type)
