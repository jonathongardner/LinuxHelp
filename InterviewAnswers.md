## Question 1
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

## Question 2
```ruby
puts special_method1('this is a test'.split('')) # => special_method1: thisisatest
puts special_method2('this is a test'.split('')) # => special_method2: this
puts special_method3('this is a test'.split('')) # => this

puts special_method1('this_is_a_test'.split('')) # => special_method1: this_is_a_test
puts special_method2('this_is_a_test'.split('')) # => special_method2: this_is_a_test
puts special_method3('this_is_a_test'.split('')) # => special_method3: this_is_a_test
```
 - Next returns the value acc and goes to the next iterable value
 - Break returns the value acc and exits the loop
 - return returns the value and exits the method

## Question 3
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

## Question 4
```ruby
options1 # => ["a", "b", {"d"=>"d"}]
options2 # => {:c=>"c", :e=>"e", :f=>"f"}
```
 - `**` only works with symbols (i.e. `{ **{ c: 'c', 'd' => 'd', e: 'e', f: 'f' } }` will return wrong argument type)
