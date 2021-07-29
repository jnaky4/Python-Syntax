import unittest


class Taco:
    tomatos = 14
    meat_chunks = 22
    sour_cream = True


myTaco = Taco()
myTaco.tomatos = 22


class Time:
    # shared attribute
    time_start = 2
    # def __init__(self):
    #     self.hours = 0
    #     self.minutes = 0

    # constructor with default arguments
    def __init__(self, hours=2, minutes=25):
        self.hours, self.minutes = self.set_time(hours, minutes)

    def set_time(self, hours, minutes):
        hours = hours + (minutes // 60)
        minutes = minutes % 60
        return hours, minutes

    def __str__(self):
        return ('{}:{}'.format(self.hours, str(self.minutes).rjust(2, '0')))


    """
    rich comparison methods:
        __lt__(self, other): less than
        __le__(self, other): less-than or equal-to
        __gt__(self, other): greater-than
        __ge__(self, other): greater-than or equal-to
        __eq__(self, other): equal to
        __ne__(self, other): not-equal to
    """
    # less than operator overloading
    def __lt__(self, other):
        return True if (self.hours * 60) + self.minutes < (other.hours * 60) + other.minutes else False

    def __gt__(self, other):
        return True if (self.hours * 60) + self.minutes > (other.hours * 60) + other.minutes else False

    def __sub__(self, other):
        return ((self.hours * 60) + self.minutes) - ((other.hours * 60) + other.minutes)

    def print_time(self):
        print(str(self.hours) + ":" + str(self.minutes).rjust(2, '0'))


my_time = Time(7, 125)
your_time = Time()

my_time.print_time()
print(my_time)


# overloading the less than operator
print(my_time < your_time)
print(my_time - your_time)

time_dict = {}
for i in range(7):
    time_dict[i] = Time(i, i)
print(time_dict[6])


# inheritance
class Item:
    def __init__(self):
        self.name = ''
        self.quantity = 0

    def set_name(self, nm):
        self.name = nm

    def set_quantity(self, qnty):
        self.quantity = qnty

    def display(self):
        print(self.name, self.quantity)


class Produce(Item):  # Derived from Item
    def __init__(self):
        Item.__init__(self)  # Call base class constructor
        self.expiration = ''

    def set_expiration(self, expir):
        self.expiration = expir

    def get_expiration(self):
        return self.expiration

item1 = Item()
item1.set_name('Smith Cereal')
item1.set_quantity(9)
item1.display()

item2 = Produce()
item2.set_name('Apples')
item2.set_quantity(40)
item2.set_expiration('May 5, 2012')
item2.display()
print('  (Expires:({}))'.format(item2.get_expiration()))


# access modifiers
# public
class Student:
    schoolName = 'XYZ School' # class attribute

    def __init__(self, name, age):
        self.name = name # instance attribute
        self.age = age # instance attribute


# protected _ leading variables
class Student1:
    _schoolName = 'XYZ School'  # protected class attribute

    def __init__(self, name, age):
        self._name = name  # protected instance attribute
        self._age = age  # protected instance attribute


# private __ leading variables
class Student3:
    __schoolName = 'XYZ School' # private class attribute

    def __init__(self, name, age):
        self.__name = name  # private instance attribute

    def __display(self):  # private method
        print('This is private method.')


# overriding
class Item:
    def __init__(self):
       self.name = ''
       self.quantity = 0

    def set_name(self, nm):
       self.name = nm

    def set_quantity(self, qnty):
       self.quantity = qnty

    def display(self):
       print(self.name, self.quantity)


class Produce(Item):  # Derived from Item
    def __init__(self):
        super().__init__()  # Call base class constructor
        # Item.__init__(self)  # does the same thing, just shows who the super class is
        self.expiration = ''

    def set_expiration(self, expir):
        self.expiration = expir

    def get_expiration(self):
        return self.expiration

    def display(self):
        print(self.name, self.quantity, end=' ')
        print('  (Expires: {})'.format(self.expiration))


item1 = Item()
item1.set_name('Smith Cereal')
item1.set_quantity(9)
item1.display()  # Will call Item's display()

item2 = Produce()
item2.set_name('Apples')
item2.set_quantity(40)
item2.set_expiration('May 5, 2012')
item2.display()  # Will call Produce's display()


"""
assertEqual(a, b)	a == b
assertNotEqual(a,b)	a != b
assertTrue(x)	bool(x) is True
assertFalse(x)	bool(x) is False
assertIs(a, b)	a is b
assertIsNot(a,b)	a is not b
assertIsNone(x)	x is None
assertIsNotNone(x)	x is not None
assertIn(a, b)	a in b
assertNotIn(a, b)	a not in b
assertAlmostEqual(a, b)	round(a - b, 7) == 0
assertGreater(a, b)	a > b
assertGreaterEqual(a, b)	a >= b
assertLess(a, b)	a < b
assertLessEqual(a, b)	a <= b
"""


# Unit testing
# User-defined class
class Circle(object):
    def __init__(self, radius):
        self.radius = radius

    def compute_area(self):
        return 3.14 * self.radius**2


# Class to test Circle
class TestCircle(unittest.TestCase):
    def test_compute_area(self):
        c = Circle(0)
        self.assertEqual(c.compute_area(), 0.0)

        c = Circle(5)
        self.assertEqual(c.compute_area(), 78.5)

    def test_will_fail(self):
        c = Circle(5)
        self.assertLess(c.compute_area(), 0)


if __name__ == "__main__":
    unittest.main()
