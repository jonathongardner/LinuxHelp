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
