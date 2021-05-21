list = [1, "HELLO"]
list2 = ["NO"]
int = list[0]
list[0] = -1
list.append(2.2)
neg = list.pop()
len(list)
list3 = list + list2

list3.index("HELLO")
list3.append(3)
print(list3.count(3))


list4 = [1,2,3,4]
min(list4)
max(list4)
sum(list4)
list.count(1)

# extends
list5 = [1,2,3,4]
list4.extend(list5)
print(list4)

# reverse a list
list4 = list4[::-1]
print(list4)

# insert before position
my_list = [5, 8]
my_list.insert(1, 1.7)

# pop removes last item
my_list = [5, 8, 14]
val = my_list.pop()
# remove index
my_list = [5, 8, 14]
val = my_list.pop(0)

# sort
my_list = [14, 5, 8]
my_list.sort()

# indexing
my_list = [5, 8, 14]
print(my_list.index(14))

# counting
my_list = [5, 8, 5, 5, 14]
print(my_list.count(5))

# iterating
nums = [1, 4, 15, 456]

max_even = None
for num in nums:
    if num % 2 == 0: # The number is even?
        if max_even == None or num > max_even:
            max_even = num

# enumerate
# User inputs string w/ numbers: '203 12 5 800 -10'
user_input = "203 12 5 800 -10"

tokens = user_input.split()  # Split into separate strings

# built in list functions
# all: true if all != 0
print(all([1, 2, 3]))
print(all([0, 1, 2]))

print(any([0, 2]))


# max
print(max([-3, 5, 25]))

# min
print(min([-3, 5, 25]))

# sum
print(sum([-3, 5, 25]))

# multi dimensional
my_list = [[10, 20], [30, 40]]
print('First nested list:', my_list[0])
print('Second nested list:', my_list[1])
print('Element 0 of first nested list:', my_list[0][0])

# enumerate
currency = [
   [1, 5, 10 ],  # US Dollars
   [0.75, 3.77, 7.53],  #Euros
   [0.65, 3.25, 6.50]  # British pounds
]
# grabs indexes as first arg in loop
for row_index, row in enumerate(currency):
   for column_index, item in enumerate(row):
       print('currency[{}][{}] is {:.2f}'.format(row_index, column_index, item))


# slicing list
boston_bruins = ['Tyler', 'Zdeno', 'Patrice']
print('Elements 0 and 1:', boston_bruins[0:2])
print('Elements 1 and 2:', boston_bruins[1:3])

# negative indexes
election_years = [1992, 1996, 2000, 2004, 2008]
print(election_years[0:-1])  # Every year except the last
print(election_years[0:-3])  # Every year except the last three
print(election_years[-3:-1])  # The third and second to last years

my_list = [5, 10, 20]
print(my_list[0:2])

my_list = [5, 10, 20, 40, 80]
print(my_list[0:5:3])

my_list = [5, 10, 20, 40, 80]
print(my_list[2:])

my_list = [5, 10, 20, 40, 80]
print(my_list[:4])

my_list = [5, 10, 20, 40, 80]
print(my_list[:])

# list iteration
my_list = [3.2, 5.0, 16.5, 12.25]

for i in range(len(my_list)):
    my_list[ i ] += 5


# list comprehension
# new_list = [expression for name in iterable]
# adding
my_list = [10, 20, 30]
list_plus_5 = [(i + 5) for i in my_list]

print('New list contains:', list_plus_5)

# convert user input to list of ints
# inp = input('Enter numbers:')
# my_list = [int(i) for i in inp.split()]
# print(my_list)

# sum each row in two dimensional list
my_list = [[5, 10, 15], [2, 3, 16], [100]]
sum_list = [sum(row) for row in my_list]
print(sum_list)

# sum min row
my_list = [[5, 10, 15], [2, 3, 16], [100]]
min_row = min([sum(row) for row in my_list])
print(min_row)

# sorting
# sort by key being max
my_list = [[25], [15, 25, 35], [10, 15]]

sorted_list = sorted(my_list, key=max)

print('Sorted list:', sorted_list)
