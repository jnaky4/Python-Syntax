import os

import plotly.graph_objects as go
import pandas as pd
import plotly.express as px
from urllib.request import urlopen
import json

import requests

from files import file_handler


def chloropleth():
    with urlopen('https://raw.githubusercontent.com/plotly/datasets/master/geojson-counties-fips.json') as response:
        counties = json.load(response)

    df = pd.read_csv("https://raw.githubusercontent.com/plotly/datasets/master/fips-unemp-16.csv",
                     dtype={"fips": str})

    print(df)
    unemployment_filepath = os.path.join(os.getcwd(), '..', 'CSV', "unemployment.csv")
    df.to_csv(unemployment_filepath)

    fig = go.Figure(go.Choroplethmapbox(geojson=counties, locations=df.fips, z=df.unemp, colorscale="Viridis",
                                        zmin=0, zmax=12, marker_opacity=0.5, marker_line_width=0))
    fig.update_layout(mapbox_style="carto-positron", mapbox_zoom=3, mapbox_center={"lat": 37.0902, "lon": -95.7129})
    fig.update_layout(margin={"r": 0, "t": 0, "l": 0, "b": 0})
    fig.show()


def mapbox_lines():
    us_cities = pd.read_csv("https://raw.githubusercontent.com/plotly/datasets/master/us-cities-top-1k.csv")

    us_cities_filepath = os.path.join(os.getcwd(), '..', 'CSV', "us_cities.csv")
    us_cities.to_csv(us_cities_filepath)

    us_cities = us_cities.query("State in ['Missouri', 'Colorado']")
    print(us_cities)
    fig = px.line_mapbox(us_cities, lat="lat", lon="lon", color="State", zoom=3, height=300)

    fig.update_layout(mapbox_style="stamen-terrain", mapbox_zoom=2.5, mapbox_center_lat=40,
                      margin={"r": 0, "t": 0, "l": 0, "b": 0})

    fig.show()


def mapbox_flights():
    # AREA EXTENT COORDINATE WGS4
    lon_min, lat_min = -125.974, 24.555
    lon_max, lat_max = -68.748, 52.214

    url_data = f"https://opensky-network.org/api/states/all?lamin={lat_min}&lomin={lon_min}&lamax={lat_max}&lomax={lon_max}"
    response = requests.get(url_data).json()
    col_name = ['icao24', 'callsign', 'origin_country', 'time_position', 'last_contact', 'long', 'lat', 'baro_altitude',
                'on_ground', 'velocity', 'true_track', 'vertical_rate', 'sensors', 'geo_altitude', 'squawk', 'spi',
                'position_source', '?']
    flight_df = pd.DataFrame(response['states'], columns=col_name)
    flight_df["squawk"] = flight_df["squawk"].fillna('0')
    print(len(flight_df["icao24"]))
    flight_df["baro_altitude"] = flight_df["baro_altitude"].fillna(0.0)
    # # flight_df = flight_df.fillna('No Data')  # replace NAN with No Data
    flight_df = flight_df.astype(dtype={
        'long': float,
        'lat': float,
        'geo_altitude': float,
        'time_position': int,
        'last_contact': int,
        'baro_altitude': float,
        'velocity': float,
        'true_track': float,
        'vertical_rate': float,
        'squawk': str,
        'spi': bool,
        'on_ground': bool
    })
    flight_df['baro_altitude'] = ((flight_df['baro_altitude'] * 3.2808) / 100).__round__(0)

    fig = px.scatter_mapbox(flight_df, lat='lat', lon='long', zoom=3, height=1000, hover_name='callsign',
                            hover_data=['squawk', 'baro_altitude'])
    fig.update_layout(mapbox_style="stamen-terrain", mapbox_zoom=2.5, mapbox_center_lat=40,
                      margin={"r": 0, "t": 0, "l": 0, "b": 0})

    fig.show()

    # while True:
    #     response = requests.get(url_data).json()
    #     flight_df = pd.DataFrame(response['states'], columns=col_name)
    #     fig.batch_update():
    #
    #     fig.show()


# mapbox_lines()
# chloropleth()
mapbox_flights()
