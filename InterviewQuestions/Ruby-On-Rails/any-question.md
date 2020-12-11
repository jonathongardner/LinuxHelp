```ruby
def some_method(x)
  puts x
  x == 2
end
```
```ruby
[some_method(1), some_method(2), some_method(3), some_method(4)].any? {|x| x }
[1, 2, 3, 4].any? {|x| some_method(x) }
```
Whats the difference?
