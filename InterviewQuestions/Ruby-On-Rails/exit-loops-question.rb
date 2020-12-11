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
