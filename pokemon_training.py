import pandas as pd

# dependencies: pandas, Pokemon.csv

# make sure the Pokemon.csv is in the same folder as this file
filename = "CSV/Pokemon.csv"



# explanation of csv reader
# https://www.delftstack.com/howto/python/python-csv-to-dictionary/
# UTF encoding error with this csv use cp1252
items = pd.read_csv(filename, index_col=0, sep=",", encoding='cp1252')

# print(items)


# https://pandas.pydata.org/docs/reference/api/pandas.DataFrame.transpose.html?highlight=transpose#pandas.DataFrame.transpose
# transpose flips the keys to be the row instead of column
pokemon_csv_data = items.transpose().to_dict(orient='series')


# example of accessing values
# key is the pokedex number, 1 is bulbasaur
# print(pokemon_dictionary)
print(pokemon_csv_data[1]['Description'])


# TODO pass in all values for a pokemon and make a pokedex class from the pokemon

class Pokedex():
    # static member
    # pokedex_dict = Pokedex.create_pokedex_dict()

    # Constructor

    def __init__(self, dexnum, level=1):
        self.dexnum = dexnum
        self.level = level
        self.name = pokemon_csv_data[dexnum]['Pokemon_Name']
        self.type1 = pokemon_csv_data[dexnum]['Type1']
        self.type2 = pokemon_csv_data[dexnum]['Type2'] if pokemon_csv_data[dexnum]['Type2'] != "-" else "Null"
        self.stage = pokemon_csv_data[dexnum]['Stage']
        self.evolve_lvl = int(pokemon_csv_data[dexnum]['Evolve_Level'])
        self.gender_ratio = pokemon_csv_data[dexnum]['Gender_Ratio']
        self.height = pokemon_csv_data[dexnum]['Height']
        self.weight = pokemon_csv_data[dexnum]['Weight']
        self.description = pokemon_csv_data[dexnum]['Description']
        self.category = pokemon_csv_data[dexnum]['Category']
        self.leveling_spd = pokemon_csv_data[dexnum]['Leveling_Speed']
        self.base_exp = pokemon_csv_data[dexnum]['Base_Exp']
        self.catch_rate = pokemon_csv_data[dexnum]['Catch_Rate']


    def time_to_evolve(self):
        # make sure the default value -1 for pokemon that dont evolve is checked
        # check if the pokemons level is high enough to evolve

        # base case
        if self.evolve_lvl != -1 and self.level < self.evolve_lvl:
            return False

        # basic intermediate and final

        # only baisc or Intermediate pokemon evolve
        if self.stage == "Basic" or self.stage == "Intermediate":
            if pokemon_csv_data[self.dexnum + 1]['Stage'] == "Intermediate" \
                    or pokemon_csv_data[self.dexnum + 1]['Stage'] == "Final":
                return True

    @staticmethod
    def create_pokedex_dict():
        pokedex_dictionary = {}
        for i in range(1, len(pokemon_csv_data)):
            pokedex_dictionary[i] = Pokedex(i)
        return pokedex_dictionary





bulbasaur = Pokedex(1,18)
bulbasaur.time_to_evolve()
print("Bulbasaur is ready to evolve?: ", bulbasaur.time_to_evolve())







#
# pokedex_entry = Pokedex(115)
# print(pokedex_entry.name)
#




def create_pokedex_dict():
    pokedex_dictionary = {}
    for i in range(1, len(pokemon_csv_data)):
        pokedex_dictionary[i] = Pokedex(i)
    return pokedex_dictionary














#

#

#


#
# Pokedex.pokedex_dictionary = Pokedex.create_pokedex_dict()
#
#
#
# # create a variable called pokedex_entry and assign a pokedex object
# pokedex_entry = Pokedex(1)
# print(pokedex_entry.name)
#
#
#
#
#
#
#
#
#
#

#
#
#
# pokedex_dictionary = create_pokedex_dict()
#
# print(pokedex_dictionary[1].name)
#
#

#
# print(type(bulbasaur.weight))