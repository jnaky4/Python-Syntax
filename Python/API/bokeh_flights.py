import os

from bokeh.plotting import figure, show
from bokeh.models import HoverTool, LabelSet, ColumnDataSource, tiles
import numpy as np
import requests
import pandas as pd
from bokeh.io import curdoc

curdoc().theme = "dark_minimal"

# AREA EXTENT COORDINATE WGS4
lon_min, lat_min = -125.974, 30.038
lon_max, lat_max = -68.748, 52.214

# <img src="https://stamen-tiles.a.ssl.fastly.net/terrain/14/2627/6331@2x.png" />
#         STAMEN_TERRAIN_RETINA='https://stamen-tiles.a.ssl.fastly.net/terrain/{Z}/{X}/{Y}@2x.png',
#         STAMEN_TERRAIN='https://stamen-tiles.a.ssl.fastly.net/terrain/{Z}/{X}/{Y}.png',

url_data = f"https://opensky-network.org/api/states/all?lamin={lat_min}&lomin={lon_min}&lamax={lat_max}&lomax={lon_max}"
response = requests.get(url_data).json()

col_name = ['icao24', 'callsign', 'origin_country', 'time_position', 'last_contact', 'long', 'lat', 'baro_altitude',
            'on_ground', 'velocity', 'true_track', 'vertical_rate', 'sensors', 'geo_altitude', 'squawk', 'spi',
            'position_source', '?']
flight_df = pd.DataFrame(response['states'], columns=col_name)

flights_df_path = os.path.join(os.getcwd(), '..', 'CSV', "flights.csv")
flight_df.to_csv(flights_df_path)

flight_df["squawk"] = flight_df["squawk"].fillna(0)
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
    'squawk': int,
    'spi': bool,
    'on_ground': bool
})
flight_df['baro_altitude'] = ((flight_df['baro_altitude'] * 3.2808) / 100).__round__(0)
# flight_df["geo_altitude"] = flight_df["geo_altitude"].astype(str).astype(float)
# def meters_to_flight_level(row):
#     row.baro_altitude = (row.baro_altitude * 3.2808) / 100
#     return row
#
#
# # print(flight_df.head())
# flight_df.apply(meters_to_flight_level, axis='columns')

flight_df['baro_altitude'] = ((flight_df['baro_altitude'] * 3.2808) / 100).__round__(0)


# print(flight_df['baro_altitude'])

# filtered = flight_df.loc[flight_df["callsign"] >= "BOE469", ['callsign', 'baro_altitude']]
#
# morefiltered = filtered.loc[filtered["callsign"] < "BOE470"]
# print(morefiltered['baro_altitude'])

# FUNCTION TO CONVERT GCS WGS84 TO WEB MERCATOR POINT
def wgs84_web_mercator_point(lon, lat):
    k = 6378137
    x = lon * (k * np.pi / 180.0)
    y = np.log(np.tan((90 + lat) * np.pi / 360.0)) * k
    return x, y


# DATA FRAME
def wgs84_to_web_mercator(df, lon="long", lat="lat"):
    k = 6378137
    df["x"] = df[lon] * (k * np.pi / 180.0)
    df["y"] = np.log(np.tan((90 + df[lat]) * np.pi / 360.0)) * k
    return df


# COORDINATE CONVERSION
xy_min = wgs84_web_mercator_point(lon_min, lat_min)
xy_max = wgs84_web_mercator_point(lon_max, lat_max)
wgs84_to_web_mercator(flight_df)
flight_df['rot_angle'] = flight_df['true_track'] * -1  # Rotation angle
icon_url = 'https://.....'  # Icon url
flight_df['url'] = icon_url

# FIGURE SETTING
# COORDINATE RANGE IN WEB MERCATOR
x_range, y_range = ([xy_min[0], xy_max[0]], [xy_min[1], xy_max[1]])
p = figure(x_range=x_range, y_range=y_range, x_axis_type='mercator', y_axis_type='mercator', sizing_mode='scale_width',
           plot_height=300)

# PLOT BASEMAP AND AIRPLANE POINTS
flight_source = ColumnDataSource(flight_df)
# tile_prov=get_provider(STAMEN_TERRAIN)
# p.add_tile()
# p.add_tile(tile_prov,level='image')
# p.image_url(url=f'https://stamen-tiles.a.ssl.fastly.net/terrain/{0}/{100}/{200}.png', x='x', y='y', source=flight_source, anchor='center', angle_units='deg', angle='rot_angle',
#             h_units='screen', w_units='screen', w=40, h=40)
p.circle('x', 'y', source=flight_source, fill_color='red', hover_color='yellow', size=10, fill_alpha=0.8, line_width=0)

# HOVER INFORMATION AND LABEL
my_hover = HoverTool()
my_hover.tooltips = [('Call sign', '@callsign'), ('Origin Country', '@origin_country'), ('velocity(m/s)', '@velocity'),
                     ('Altitude(fl)', '@baro_altitude')]
labels = LabelSet(x='x', y='y', text='callsign', level='glyph',
                  x_offset=5, y_offset=5, source=flight_source, render_mode='canvas',
                  text_font_size="8pt", text_color='white')
p.add_tools(my_hover)
p.add_layout(labels)

show(p)
