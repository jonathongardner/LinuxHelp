import random
import time

class GuessMyNumber:
  def __init__(self, max_number=100, mock_random=None):
    self.max_number = max_number
    self.number = mock_random or random.randint(1, max_number)

  def guess(self, guess):
    if guess == self.number:
      return 0
    elif guess > self.number:
      print(f"{guess} ({self.number}) To high, waiting 60 seconds")
      time.sleep(60)
      return 1
    print(f"{guess} ({self.number}) To low, waiting 1 seconds")
    time.sleep(1)
    return -1

# write a method that will guess the number for max 100 in under 100 seconds
def guess_my_number():
  gmm = GuessMyNumber(mock_random=99)
  guess = 10
  while True:
    valid = gmm.guess(guess)
    if valid == 0:
      return guess
    elif valid == 1:
      guess -= 9
      break
    guess += 10

  while True:
    valid = gmm.guess(guess)
    if valid == 0:
      return guess

    guess += 1



start = time.time()
print(guess_my_number())
end = time.time()
print(f"Time taken: {end - start}")
