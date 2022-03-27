import requests

# url = "https://globe.adsbexchange.com/data/globe_6235.binCraft"

url = "https://foaas.com/version"
headers = {"Accept": "application/json"}
response = requests.get(url, headers=headers)

print(response)
print(response.text)
