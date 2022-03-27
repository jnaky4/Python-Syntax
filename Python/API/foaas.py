import requests
import pprint
import pandas as pd

url = "https://foaas.com/bag/jake"
headers = {"Accept": "application/json"}
response = requests.get(url, headers=headers)

val = response.text
# print(type(val))
print(response.text)

pd.set_option('display.width', None)
pd.set_option('display.max_rows', None)
pd.set_option('display.max_columns', None)


json = pd.read_json(val)
print(json)
pprint.pprint(json)
print(response.text)