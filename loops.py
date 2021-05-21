"""
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

# range
origins = [4, 8, 10]

for index in range(len(origins)):
    value = origins[index]  # Retrieve value of element in list.
    print('Element {}: {}'.format(index, value))

# list.index()
origins = [4, 8, 10]

for value in origins:
    index = origins.index(value)  # Retrieve index of value in list
    print('Element {}: {}'.format(index, value))

# enumerate
origins = [4, 8, 10]

for (index, value) in enumerate(origins):
    print('Element {}: {}'.format(index, value))