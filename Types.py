from typing import List, Set, Dict, Tuple, Optional, Callable, Iterator, Union, Optional

a: int = 1
b: float = 1.0
c: bool = True
d: str = "test"
e: bytes = b"test"
f: List[int] = [1]
g: Set[int] = {6, 7}

# for Python 3.9 and Later
"""
# For collections, the type of the collection item is in brackets
# (Python 3.9+)
x: list[int] = [1]
x: set[int] = {6, 7}
"""

# For mappings, we need the types of both keys and values
# h: dict[str, float] = {'field': 2.0}  # Python 3.9+
h: Dict[str, float] = {'field': 2.0}

# For tuples of fixed size, we specify the types of all the elements
# i: tuple[int, str, float] = (3, "yes", 7.5)  # Python 3.9+
i: Tuple[int, str, float] = (3, "yes", 7.5)

# For tuples of variable size, we use one type and ellipsis
# j: tuple[int, ...] = (1, 2, 3)  # Python 3.9+
j: Tuple[int, ...] = (1, 2, 3)

# Use Optional[] for values that could be None
# x: Optional[str] = some_function()
# Mypy understands a value can't be None in an if-statement
# if x is not None:
#     print(x.upper())
# # If a value can never be None due to some invariants, use an assert
# assert x is not None
# print(x.upper())


# This is how you annotate a function definition
def stringify(num: int) -> str:
    return str(num)


# And here's how you specify multiple arguments
def plus(num1: int, num2: int) -> int:
    return num1 + num2


# Add default value for an argument after the type annotation
def fa(num1: int, my_float: float = 3.5) -> float:
    return num1 + my_float


# This is how you annotate a callable (function) value
x: Callable[[int, float], float] = fa


# A generator function that yields ints is secretly just a function that
# returns an iterator of ints, so that's how we annotate it
def ga(n: int) -> Iterator[int]:
    p = 0
    while p < n:
        yield p
        p += 1


# You can of course split a function annotation over multiple lines
def send_email(address: Union[str, List[str]],
               sender: str,
               cc: Optional[List[str]],
               bcc: Optional[List[str]],
               subject='',
               body: Optional[List[str]] = None
               ) -> bool:
    ...


# An argument can be declared positional-only by giving it a name
# starting with two underscores:
def quux(__x: int) -> None:
    pass


quux(3)  # Fine
quux(__x=3)  # Error


class Solution(object):
    def __init__(self):
        self.x = 4

    def beautiful_array(self, n: int) -> List[int]:
        list = []
        for z in range(n):
            list.append(z)
        return list


s = Solution()

print(s.beautiful_array(4))
