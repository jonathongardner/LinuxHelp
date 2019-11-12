## Question 1
```ruby
my_hash = { key1: 'value1', 'key2' => 'value2' }
```
What would `my_hash['key1']` return?

## Question 2
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

## Question 3
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

## Bonus
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
