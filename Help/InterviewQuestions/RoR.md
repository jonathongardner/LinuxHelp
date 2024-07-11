# RoR interview questions
  - How can you manage multiple versions of ruby?
    - What if you have multiple ruby projects with difference gems?
  - How can you login into a remote computer/virtual machine?
  - How can you search for a word/phrase in multiple files/directories in a linux system?
  - What git workflows have you used?


## Coding Questions
```ruby
class Post < ActiveRecord::Base
  has_one :post_metadata
  # fields:
  # message: text
  # type: string
  # owner: string
  # ...
  before_validation :initialize_to
  validates :message, :to, :owner, presence: true

  def initialize_to
    to ||= 'everyone'
  end
end

class PostMetadata < ApplicationRecord
  belongs_to :post
  # fields:
  # likes_count: bigint
  # post_id: bigint fk to post.id
  # ...
end

class User < ApplicationRecord
  # fields:
  # username: string
  # ...

  # Gets total number of likes for all post by a given user
  def get_total_likes(username)
    likes = 0
    mega_likes = false
    Post.where(owner: username).each do |p|
      likes += p.post_metadata.likes_count
    end
    if likes > 100_000
      mega_likes = true
    end
    {
      likes: likes,
      mega_likes: mega_likes
    }
  end
end

class SomeController < ApplicationController
  #...
  def set_like_information
    @like_information = current_user.get_total_likes(current_user.username)
  end
  #...
end
```
Questions:
  - Whats wrong with `initialize_type`, how could we debug this?
  - Whats wrong with `set_like_information` (and the methods it uses)/How would you refactor them?

## More coding
```ruby
Post.includes(:post_metadata).each do |post|
  puts "#{message} by #{owner} has #{...} likes"
end
Post.eager_load(:post_metadata).each do |post|
  puts "#{message} by #{owner} has #{...} likes"
end
Post.preload(:post_metadata).each do |post|
  puts "#{message} by #{owner} has #{...} likes"
end
Post.joins(:post_metadata).pluck(:owner, :message, :likes_count) do |data|
  puts "#{message} by #{owner} has #{...} likes"
end
Post.left_joins(:post_metadata).pluck(:owner, :message, :likes_count) do |data|
  puts "#{message} by #{owner} has #{...} likes"
end
```
Questions:
  - What is the differences for each?
  - How can we make this work for each example?

## Docker Questions
```Dockerfile
FROM ruby:3.0-alpine
WORKDIR /src/interview
COPY . .
RUN bundle install
```
Questions:
  - How can we make this better?
