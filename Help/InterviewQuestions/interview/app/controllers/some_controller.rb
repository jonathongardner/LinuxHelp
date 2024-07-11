class SomeController < ApplicationController
  #...
  def set_like_information
    @like_information = current_user.get_total_likes(current_user.username)
  end
  #...
end
