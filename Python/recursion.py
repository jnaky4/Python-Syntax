def count_down(count):
    if count == 0:
        print('Go!')
    else:
        print(count)
        count_down(count - 1)


count_down(2)

# searching sored list
def find(lst, item, low, high):
    """
    Finds index of string in list of strings, else -1.
    Searches only the index range low to high
    Note: Upper/Lower case characters matter
    """
    range_size = (high - low) + 1
    mid = (high + low) // 2

    if item == lst[mid]:  # Base case 1: Found at mid
        pos = mid
    elif range_size == 1:  # Base case 2: Not found
        pos = -1
    else:  # Recursive search: Search lower or upper half
        if item < lst[mid]:  # Search lower half
            pos = find(lst, item, low, mid)
        else:  # Search upper half
            pos = find(lst, item, mid+1, high)

    return pos

attendees = []

attendees.append('Adams, Mary')
attendees.append('Carver, Michael')
attendees.append('Domer, Hugo')
attendees.append('Fredericks, Carlo')
attendees.append('Li, Jie')

name = "Domer, Hugo"
pos = find(attendees, name, 0, len(attendees)-1)

if pos >= 0:
    print('Found at position {}.'.format(pos))
else:
    print('Not found.')

# factiorial recursion
def nfact(n):
    if n == 1 or n == 0:  # Base case
        fact = 1
    else:  # Recursive case
        fact = n * nfact(n - 1)
    return fact

# Get n from user, print nfact(n)


def fibonacci(v1, v2, run_cnt):
    print(v1, '+', v2, '=', v1+v2)

    if run_cnt <= 1:  # Base case:
                      # Ran for user's number of steps
        pass  # Do nothing
    else:             # Recursive case
        fibonacci(v2, v1+v2, run_cnt-1)


print ('This program outputs the\n'
       'Fibonacci sequence step-by-step,\n'
       'starting after the first 0 and 1.\n')

run_for = int(input('How many steps would you like?'))

fibonacci(0, 1, run_for)


def gcd(n1, n2):
    if n1 % n2 == 0:           # n2 is a common factor
        return n2
    else:
        return gcd(n2,n1%n2)


print ('This program outputs the greatest '
       'common divisor of two numbers.\n')

num1 = int(input('Enter first number: '))
num2 = int(input('Enter second number: '))

if (num1 < 1) or (num2 < 1):
    print('Note: Neither value can be below 1.')
else:
    my_gcd = gcd(num1, num2)
    print('Greatest common divisor =', my_gcd)