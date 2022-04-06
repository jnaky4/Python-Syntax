import time
import pandas as pd

# initalize dictionary to specific size
a = dict.fromkeys((range(4000000)))
print(len(a))

player1 = {}  # empty dictionary

players = {
    'jake': 10,
    'john': 7
}
# add a item to dictionary
players['jack'] = 4
players["jake"] = 69

# print(players)
# print(players["jake"])

del players["jack"]

# clear a dictionary
my_dict = {'Ahmad': 1, 'Jane': 42}
my_dict.clear()
# print(my_dict)

# get key from dictionary, if doesnt exist returns N/A
my_dict = {'Ahmad': 1, 'Jane': 42}
# print(my_dict.get('Jane', 'N/A'))
# print(my_dict.get('Chad', 'N/A'))

# merging dictionaries, if duplicate keys, the dict being merged into keys are replaced
my_dict = {'Ahmad': 1, 'Jane': 42}
my_dict.update({'John': 50})
# print(my_dict)

# removes a key
my_dict = {'Ahmad': 1, 'Jane': 42}
val = my_dict.pop('Ahmad')
# print(my_dict)

# iterations
# iterate items
num_calories = dict(Coke=90, Coke_zero=0, Pepsi=94)
for soda, calories in num_calories.items():
    print('{}: {}'.format(soda, calories))

# iterate keys
num_calories = dict(Coke=90, Coke_zero=0, Pepsi=94)
for soda in num_calories.keys():
    print(soda)

# iterate values
num_calories = dict(Coke=90, Coke_zero=0, Pepsi=94)
for soda in num_calories.values():
    print(soda)

# list conversion
solar_distances = dict(mars=219.7e6, venus=116.4e6, jupiter=546e6, pluto=2.95e9)
list_of_distances = list(solar_distances.values())  # Convert view to list

# sort new list
sorted_distance_list = sorted(list_of_distances)
closest = sorted_distance_list[0]
next_closest = sorted_distance_list[1]

print('Closest planet is {:.4e}'.format(closest))
print('Second closest planet is {:.4e}'.format(next_closest))

# nested dictionaries
students = {}
students['Jose'] = {'Grade': 'A+', 'StudentID': 22321}

print('Jose:')
print(' Grade: {}'.format(students['Jose']['Grade']))
print(' ID: {}'.format(students['Jose']['StudentID']))

grades = {
    'John Ponting': {
        'Homeworks': [79, 80, 74],
        'Midterm': 85,
        'Final': 92
    },
    'Jacques Kallis': {
        'Homeworks': [90, 92, 65],
        'Midterm': 87,
        'Final': 75
    },
    'Ricky Bobby': {
        'Homeworks': [50, 52, 78],
        'Midterm': 40,
        'Final': 65
    },
}

# Nested Dictionary example
user_input = "Ricky Bobby"

while user_input != 'exit':
    if user_input in grades:
        # Get values from nested dict
        homeworks = grades[user_input]['Homeworks']
        midterm = grades[user_input]['Midterm']
        final = grades[user_input]['Final']

        # print info
        for hw, score in enumerate(homeworks):
            print('Homework {}: {}'.format(hw, score))

        print('Midterm: {}'.format(midterm))
        print('Final: {}'.format(final))

        # Compute student total score
        total_points = sum([i for i in homeworks]) + midterm + final
        print('Final percentage: {:.1f}%'.format(100 * (total_points / 500.0)))

    user_input = "exit"

gmt = time.gmtime()  # Get current Greenwich Mean Time

print('Time is: {:02d}/{:02d}/{:04d}  {:02d}:{:02d} {:02d} sec' \
      .format(gmt.tm_mon, gmt.tm_mday, gmt.tm_year, gmt.tm_hour, gmt.tm_min, gmt.tm_sec))

gmt = time.gmtime()  # Get current Greenwich Mean Time

print('Time is: %(month)02d/%(day)02d/%(year)04d  %(hour)02d:%(min)02d %(sec)02d sec' % \
      {
          'year': gmt.tm_year, 'month': gmt.tm_mon, 'day': gmt.tm_mday,
          'hour': gmt.tm_hour, 'min': gmt.tm_min, 'sec': gmt.tm_sec
      }
      )

dict = {
    # equipment_id
    45: {
        # company_id
        1: {
            "eq_name": "leaf blower",
            "comp_name": "Riven-dell computers",
            "day": 20,
            "week": 100,
            "month": 400
        },
        2: {
            "eq_name": "leaf blower",
            "comp_name": "Abalonga Wonga",
            "day": 40,
            "week": 350,
            "month": 4000
        }
    },
    22: {

        # company_id
        1: {
            "eq_name": "tacos",
            "comp_name": "six flags",
            "day": 20,
            "week": 100,
            "month": 400
        },
        2: {
            "eq_name": "chainsaw 3000",
            "comp_name": "Abalonga Wonga",
            "day": 40,
            "week": 350,
            "month": 4000
        }
    }
}

# Nested Dictionary Example
javascript_dict = {}

for key, row in dict.items():
    key_name = key
    for key2, column in row.items():
        print(key2)
        print(column)

eq_id = 22
comp_id = 1
# print(dict[eq_id][comp_id])


"""
Nested dictionary Creation

"""
# Create Nested Dictionary
listy = [1, 2, 3, 4]
layer2 = [6, 54, 3]
test_dict = {}

# first layer creation
for item in listy:
    if item not in test_dict:
        test_dict[item] = {}

# second layer
for item in test_dict:
    for li in layer2:
        if li not in test_dict[item]:
            test_dict[item][li] = layer2
print(test_dict)

"""
CSV to Dictionary
"""

# csv to dictionary example
filename = "CSV/Pokemon.csv"

# explanation of csv reader
# https://www.delftstack.com/howto/python/python-csv-to-dictionary/
# UTF encoding error with this csv use cp1252
items = pd.read_csv(filename, index_col=0, sep=",", encoding='cp1252')

# https://pandas.pydata.org/docs/reference/api/pandas.DataFrame.transpose.html?highlight=transpose#pandas.DataFrame.transpose
# transpose flips the keys to be the row instead of column
pokemon_csv_data = items.transpose().to_dict(orient='series')

# example of accessing values
# key is the pokedex number, 1 is bulbasaur
# print(pokemon_dictionary)
print(pokemon_csv_data[1]['Description'])
