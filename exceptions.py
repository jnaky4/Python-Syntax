try:
    result = 5 / 0

# only 1 exception will be caught in a try block
except Exception as e:
    print(e)
    try:
        raise TypeError
    except TypeError:
        print("Caught")
except ZeroDivisionError:
    print("Zero")
except (ZeroDivisionError, TypeError):
    print("reached")

finally:
    print("do this regardless")

# Define a custom exception type
class LessThanZeroError(Exception):
    def __init__(self, value):
        self.value = value
