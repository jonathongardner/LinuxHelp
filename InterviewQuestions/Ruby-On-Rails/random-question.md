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
