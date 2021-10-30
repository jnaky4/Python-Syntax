# set is unordered collection, or unique elements
# accepts sequence of iterable objects
# mutable

# Create a set using the set() function.
nums1 = set([1, 2, 3])

# Create a set using a set literal.
nums2 = { 9, 8, 7 }
print(sorted(nums2))

len(nums1)
# adds elements from one set to another
# update sorts set
nums1.update(nums2)

print(nums1)

nums1.add(11)
nums1.remove(11)
# removes a random element from the list
nums2.pop()
set.clear(nums2)



"""A set is often used to reduce a list of items that potentially contains duplicates into a collection of unique values.
Simply passing a list into set() will cause any duplicates to be omitted in the created set."""
# Initial list contains some duplicate values
first_names = ['Harry', 'Hermione', 'Ron', 'Harry', 'Albus', 'Ron', 'Ron']

# Creating a set removes any duplicate values
names_set = set(first_names)
first_names = list(names_set)
print(first_names)


# Set Operations

# Create sets
names1 = {'Corrin', 'Pedro', 'Khan', 'Dean'}
names2 = {'Emilia', 'Kara', 'Corrin', 'Tia'}
names3 = {'Karat', 'Kara', 'Karah', 'Khan'}
names4 = {'Khan', 'Dean', 'Pascale'}

# Union names1 and names2
result_set = names1.union(names2)

# Intersection btwn result_set and names3
result_set = result_set.intersection(names3)

# Difference btwn result_set and names4
result_set = result_set.difference(names4)

