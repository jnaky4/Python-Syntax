import requests

url = "https://globe.adsbexchange.com/data/globe_6235.binCraft"

response = requests.get(url)
print(response)
