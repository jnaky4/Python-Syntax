from collections import namedtuple

# immutable
tup = (2,3,4)
len(tup)
min(tup)
max(tup)
tup.count(2)
print(tup[0])

# named tuple factory function
# creates class definition
Car = namedtuple('Car', ['make','model','price','horsepower','seats'])
chevy_blazer = Car('Chevrolet', 'Blazer', 32000, 275, 8)  # Use the named tuple to describe a car
chevy_impala = Car('Chevrolet', 'Impala', 37495, 305, 5)  # Use the named tuple to describe a different car

print(chevy_blazer)
print(chevy_impala)
