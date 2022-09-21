import requests
import pprint
import pandas as pd
from io import StringIO

url = "https://pokeapi.co/api/v2/pokemon/bulbasaur"
headers = {"Accept": "application/json"}
response = requests.get(url, headers=headers)

val = response.text
# print(type(val))
# print(response.text)

# pd.set_option('display.width', None)
# pd.set_option('display.max_rows', None)
# pd.set_option('display.max_columns', None)
#
#
# json = pd.read_json(StringIO(val))
# print(json)
# pprint.pprint(response.text)
