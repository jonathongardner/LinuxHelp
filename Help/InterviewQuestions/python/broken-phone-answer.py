# write a method that will guess the number for max 100 in under 100 seconds
def fast_more_tries(batch_size=10):
  gmm = GuessMyNumber()
  guess = batch_size
  while True:
    valid = gmm.guess(guess)
    if valid == 0:
      return guess
    elif valid == 1: # to high
      guess -= (batch_size - 1)
      break
    guess += batch_size

  while True:
    valid = gmm.guess(guess)
    if valid == 0:
      return guess

    guess += 1

def slow_few_tries():
  gmm = GuessMyNumber()
  def rec(start, end):
    guess = start + math.ceil((end - start) / 2)
    valid = gmm.guess(guess)
    if valid == 0:
      return
    elif valid == 1: # to high
      rec(start, guess)
    else:
      rec(guess, end)

  rec(1, 100)