```ruby
[some_method(1), some_method(2), some_method(3), some_method(4)].any? {|x| x }
1
2
3
4
```
```ruby
[1, 2, 3, 4].any? {|x| some_method(x) }
1
2
```
The first method is evaluated and than the `any?` operator is run, the second evaluates the method in the `any?` so it short circuits.
