import requests
import pandas as pd
import time
import csv

#AREA EXTENT COORDINATE WGS4
lon_min, lat_min = -125.974, 30.038
lon_max, lat_max = -68.748, 52.214

# real time flight data within the US
url_data = f"https://opensky-network.org/api/states/all?lamin={lat_min}&lomin={lon_min}&lamax={lat_max}&lomax={lon_max}"
print(url_data)

csv_data = "./CSV/flight_data.csv"

response = requests.get(url_data).json()

# #LOAD TO PANDAS DATAFRAME
col_name = ['icao24','callsign','origin_country','time_position','last_contact','long','lat','baro_altitude','on_ground','velocity',
'true_track','vertical_rate','sensors','geo_altitude','squawk','spi','position_source','?']
flight_df = pd.DataFrame(response['states'], columns=col_name)
flight_df = flight_df.fillna('No Data') #replace NAN with No Data

#REQUEST INTERVAL
sleep_time = 10

while col_name != '':
    with open(csv_data, 'w') as csv_file:
        csv_writer = csv.writer(csv_file, delimiter=',', quotechar='"', quoting=csv.QUOTE_ALL)
        csv_writer.writerow(col_name)
        response = requests.get(url_data).json()
        try:
            n_response = len(response['states'])
        except Exception:
            pass
        else:
            for i in range(n_response):
                info = response['states'][i]
                csv_writer.writerow(info)
    time.sleep(sleep_time)
    print('Get', len(response['states']), 'data')




# # get top 5 rows
# print(flight_df.head())
#
# # size of dataFrame
# print(flight_df.shape)
#
# # read csv
# grades = pd.read_csv(csv_data, index_col=0)
#
# # filter by column name
# print(flight_df["callsign"])
# # specific row
# print(flight_df["callsign"][0])
