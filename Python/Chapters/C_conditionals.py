a = True
b = False


print(a and b)
print(a or b)
print(not a)

#nor
# true if both false
print(not (a or b))
#nand
# true if either false
print(not (a and b))

# conditional expression
my_var = "hello" if 4 > 2 else "bye"
print(my_var)