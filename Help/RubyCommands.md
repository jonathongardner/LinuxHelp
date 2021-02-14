## RVM
#### Create gemset
```BASH
rvm gemset create $gemset_name
```
#### Create alias
```BASH
rvm alias create $alias_name $ruby_version@$gemset_name
```

## Build Gems
```BASH
gem build awesome_gem.gemspec
```
