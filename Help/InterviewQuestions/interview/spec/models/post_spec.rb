require "rails_helper"

RSpec.describe 'Post' do
  it "initialized 'to' to 'everyone' after save" do
    post = Post.create(message: 'Nice hat.', owner: 'foo')
    expect(post.save).to(be_truthy, "Didnt save post #{post.errors.full_messages}")
    expect(post.to).to eq('everyone')
  end
  it "it used 'to' if passed" do
    post = Post.create(message: 'Nice hat.', to: 'bar', owner: 'foo')
    expect(post.save).to(be_truthy, "Didnt save post #{post.errors.full_messages}")
    expect(post.to).to eq('bar')
  end
end
