import random

class GuessMyNumber:
  def __init__(self, max_number=100, mock_random=None):
    self.max_number = max_number
    self.number = mock_random or random.randint(1, max_number)
    self.tries = 0
    self.time_sum = 0

  def guess(self, guess):
    self.tries += 1
    if guess == self.number:
      print(f"{guess} ({self.number}) Correct, time taken: {self.time_sum}, tries: {self.tries}")
      return 0
    elif guess > self.number:
      print(f"{guess} ({self.number}) To high, waiting 60 seconds")
      self.time_sum += 60
      return 1
    print(f"{guess} ({self.number}) To low, waiting 1 seconds")
    self.time_sum += 1
    return -1

# write a method that will guess the number for max 100 in under 100 seconds
def guess_my_number():
  gmm = GuessMyNumber()

guess_my_number()

