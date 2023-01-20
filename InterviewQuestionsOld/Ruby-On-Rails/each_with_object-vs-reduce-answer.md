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
