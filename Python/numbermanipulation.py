import math
from typing import List


class NumberManipulation:
    # count number of digits in a number
    @staticmethod
    def count_digits(x: int) -> int:
        if x == 0:
            digits = 1
        else:
            digits = int(math.log10(x)) + 1
        return digits

    # reverse a number unless greater than 32bit int
    @staticmethod
    def reverse_num(x: int) -> int:
        a = 0
        if x > 0:  # handle positive numbers
            a = int(str(x)[::-1])
        if x <= 0:  # handle negative numbers
            a = -1 * int(str(x * -1)[::-1])
        # handle 32 bit overflow
        mina = -2 ** 31
        maxa = 2 ** 31 - 1
        if a not in range(mina, maxa):
            return 0
        else:
            return a

    @staticmethod
    def reverse_2(x: int) -> int:
        if x >= 2 ** 31 - 1 or x <= -2 ** 31:
            return 0
        else:
            strg = str(x)
            if x >= 0:
                revst = strg[::-1]
            else:
                temp = strg[1:]
                temp2 = temp[::-1]
                revst = "-" + temp2
            if int(revst) >= 2 ** 31 - 1 or int(revst) <= -2 ** 31:
                return 0
            else:
                return int(revst)

    # returns true if number is a palindrome
    @staticmethod
    def is_palindrome(x: int) -> bool:
        return str(x) == (str(x)[::-1])

    """
    Given an array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.

    Return the indices of the two numbers (1-indexed) as an integer array answer of size 2, where 1 <= answer[0] < answer[1] <= numbers.length.

    The tests are generated such that there is exactly one solution. You may not use the same element twice.
    """
    @staticmethod
    def twoSum(numbers: List[int], target: int) -> List[int]:
        s = {}

        for i, n in enumerate(numbers):
            comp = target - n
            if comp in s:
                return [s[comp] + 1, i + 1]
            s[n] = i


print(NumberManipulation.count_digits(1234))

print(NumberManipulation.is_palindrome(12321))
