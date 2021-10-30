#TODO add google translate

# slicing
url = 'http://en.wikipedia.org/wiki/Turing'
domain = url[7:23]  # Read 'en.wikipedia.org' from url
print(domain)

# to all but last char
print(url[:-1])

# reverse
print(url[::-1])

# slice stride, print every other character
print(url[::2])


# formatting example
format_string = '{name:16}{goals:8}'
print(format_string.format(name='Player Name'.rjust(10, " "), goals='Goals'.rjust(10, " ")))
print('-' * 24)
print(format_string.format(name='Sadio Mane'.rjust(10, " "), goals=22))
print(format_string.format(name='Gabriel Jesus'.rjust(10, " "), goals=7.))


# replace
phrase = 'One day I will have three goats, six horses, and nine llamas.'
phrase = phrase.replace('one', 'uno')
phrase = phrase.replace('three', 'tres')
phrase = phrase.replace('six', 'seis')
phrase = phrase.replace('nine', 'nueve')
print('Translation:', phrase)


# find
my_str = "Boo Hoo!"
print(my_str.find('!'))
# start at index
my_str.find('oo', 2)
# start, end
my_str.find('oo', 2, 4)
# reverse find
my_str.rfind("!")
#count
my_str.count('oo')


superhero_name = "Superman batman"
# in
if 'batman' in superhero_name:
    print(True)

# comparison
print("Hello" == "Hello")
# compares char by char, Y > A
print('Yankee Sierra' > 'Amy Wise')
# in
print('seph' in 'Joseph')
# is identity operator
# jake = "jake"
# print("jake" is jake)
# print("Jake" is jake)

"""
String comparison
isalnum() -- Returns True if all characters in the string are lowercase or uppercase letters, or the numbers 0-9.
isdigit() -- Returns True if all characters are the numbers 0-9.
islower() -- Returns True if all cased characters are lowercase letters.
isupper() -- Return True if all cased characters are uppercase letters.
isspace() -- Return True if all characters are whitespace.
startswith(x) -- Return True if the string starts with x.
endswith(x) -- Return True if the string ends with x.

New Strings
capitalize() -- Returns a copy of the string with the first character capitalized and the rest lowercased.
lower() -- Returns a copy of the string with all characters lowercased.
upper() -- Returns a copy of the string with all characters uppercased.
strip() -- Returns a copy of the string with leading and trailing whitespace removed.
title() -- Returns a copy of the string as a title, with first letters of words capitalized.
"""

# split
string = 'Music/artist/song.mp3'
my_tokens = string.split('/')
print(my_tokens)
string = 'I love python'
my_tokens = string.split()
print(my_tokens)

# join
web_path = [ 'www.website.com', 'profile', 'settings' ]
separator = '/'
url = separator.join(web_path)
print(url)

list = ["Jake", "Taylor", "Sam", "Julia","Jimmy"]
print(",".join(list))

#
path = "C:/Users/Wolfman/Documents/report.pdf"
new_separator = '\\'
tokens = path.split('/')
print(new_separator.join(tokens))

