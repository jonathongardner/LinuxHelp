require "test_helper"

class PostTest < ActiveSupport::TestCase
  test "that for is initialized to 'everyone' after save" do
    post = Post.create(message: 'Nice hat.', owner: 'foo')
    assert post.save, "Didnt save post #{post.errors.full_messages}"
    assert_equal 'everyone', post.to, 'should default to general'
  end
  test "that to is used if passed" do
    post = Post.create(message: 'Nice hat.', to: 'bar', owner: 'foo')
    assert post.save, "Didnt save post #{post.errors.full_messages}"
    assert_equal 'bar', post.to, 'should default to general'
  end
end
