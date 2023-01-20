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
