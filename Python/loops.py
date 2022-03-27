
f"""
Range	Generated sequence	Explanation
range(5)
0 1 2 3 4
Every integer from 0 to 4.

range(0, 5)
0 1 2 3 4
Every integer from 0 to 4.

range(3, 7)
3 4 5 6
Every integer from 3 to 6.

range(10, 13)
10 11 12
Every integer from 10 to 12.

range(0, 5, 1)
0 1 2 3 4
Every 1 integer from 0 to 4.

range(0, 5, 2)
0 2 4
Every 2nd integer from 0 to 4.

range(5, 0, -1)
5 4 3 2 1
Every 1 integer from 5 down to 1

range(5, 0, -2)
5 3 1
Every 2nd integer from 5 down to 1
"""

# start here
# while loop: (condition true)
# counter = 1
# while counter < 6:
#     print(counter)
#     counter += 1
# else:
#     print("i is no longer less than 6")
# print()


# range: count to range 10, 0-9
for numbers in range(10):
    print(numbers, end=" ")  # end:" " print on the same line
print()

# List comprehension: better way of doing that
[print(number, end=" ") for number in range(5)]  # [f(n) for n in range(len(L))]
print()

# # range over the length of a list
# list_of_numbers = [1, 2, 2, 3, -1]
# for index in range(len(list_of_numbers)):
#     number = list_of_numbers[index]  # Retrieve value of element in list.
#     print(f"""Element {index}: {number}""", end=" ")  # Better way of printing
# print()

list_of_items = [10, "taco", .22, True, b"taco", [1, 2, 3], (1, 2, 2, 4), {'field': 2.0}]  # list of items

for item in list_of_items:
    print(item, end=" ")
print()


[print(item) for item in list_of_items]

# could be
[print(item, end=" ")for item in list_of_items]

# but becomes
[print(f"[{list_of_items.index(item)}]={item}", end=" ")for item in list_of_items]
print()

# uhh
[
    print(f"[{list_of_items.index(item)}]={item}", end=" ")
    for item in list_of_items
]
print()

# [print(x) while x > list_of_items]

# enumerate:
# list_of_items = [10, "taco", .22, True, b"taco", [1, 2, 3], (1, 2, 2, 4), {'field': 2.0}]

# for (index, item) in enumerate(list_of_items):
#     print(f"""Enumerated {index}: {item}""")
# print()

# use the typing library to specify types
from typing import List, Set, Dict

# # variable types
# a: int = 1
# b: float = 1.0
# c: bool = True
# d: str = "test"
# e: bytes = b"test"
# f: List[int] = [1, 2, 3]
# g: Set[int] = {6, 7, 7, 6}
# h: Dict[str, float] = {'field': 2.0}
#
# list_of_numbers: List[int] = [1, 2, 3, 8, -1]

# for (index, number) in enumerate(list_of_numbers):
#     print(f"""Element {index}: {number}""")
# print()


