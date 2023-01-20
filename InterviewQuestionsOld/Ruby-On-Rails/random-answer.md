```ruby
options1 # => ["a", "b", {"d"=>"d"}]
options2 # => {:c=>"c", :e=>"e", :f=>"f"}
```
 - `**` only works with symbols (i.e. `{ **{ c: 'c', 'd' => 'd', e: 'e', f: 'f' } }` will return wrong argument type)
