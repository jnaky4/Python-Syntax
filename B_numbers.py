# make sure you understand variables before you start this section
# if you want to see what a line is doing, put the line inside a print()


variable = 2
number = 5
aFloatNumber = 2.224

number = variable # you can assign one variable to another
# number is now 2

# Examples of what you can do
# add 1 to the number stored in number its now 3
number = number + 1
# short hand version
number += 1

# multiplication
number = number * number
number *= number

# division
number = number / 2
number /= number

# what if number is 5 / 2 = 2.5 its not an integer anymore
# Python will automatically change it to a float 2.5
# if you dont want this to happen, you can do integer division
number = 5 // 2 # this will round the number down to 2 and keep it an integer

"""
    Learn this it is HANDY!
    % is called the modulus operator and it give the remainder after division
"""
remainder = 10 % 3 # 10 / 3 = 3 with a remainder of 1, remainder will get assigned

