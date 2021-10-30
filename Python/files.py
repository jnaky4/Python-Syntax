import os
import struct
import pandas as pd
import struct, csv, pickle

#TODO add Pickle


# Open a file regardless of OS
# grabs route correctly independent of OS routing:
#   Linux/Mac: CSV//Pokemon.csv
#   Windows: CSV\\Pokemon.csv
pokemon_csv = os.path.join('CSV', "Pokemon.csv")
base_stats_csv = os.path.join('CSV', "Base_Stats.csv")


# with: opens file, execute block and automatically close
with open("test.txt", 'r') as f:
    print(f.readline())



# using Pandas Library to read csv
filename = "test.csv"
items = pd.read_csv(filename).to_dict("records")
print(items)



# read
f_read = open("test.txt", "r")
print(f_read)
print(f_read.readline())
f_read.close()

# write
f_write = open("test.txt", "w")
f_write.write("Overwriting Contents")
f_write.close()

# append
f_append = open("test.txt", "a")
f_append.write("\nAdd me some letters!")
f_append.close()

# create empty file
# f_empty_file = open("test2.txt", "x")


# if os.path.exists("test2.txt"):
#     os.remove("test2.txt")

# wrie output buffer to disk every 100 bytes
f = open('myfile.txt', 'w', buffering=100)
# force output buffer to be written to disk
f.flush()
# ensure that all internal buffers associated with f are written to disk
os.fsync(f.fileno())

f.close()

# Hardcoded not recommended
path = "C:\\Users\\Hyperlight Drifter\\Documents\\GitHub\\Pokemon-Remaster\\Assets\\Resources\\Images\\PokemonImages\\002Ivysaur.png"


print('Size of file:', os.path.getsize(path), 'bytes')

# Windows requires \\ with full path
p = os.path.join('C:\\', 'Users', 'Hyperlight Drifter', "Documents", "Github")
if os.path.exists(p):
    print(p)

# splitting hard coded path into list
tokens = 'C:\\Users\\BWayne\\tax_return.txt'.split(os.path.sep)
print(tokens)

# splitting head and tail
p = os.path.join('C:\\', 'Users', 'BWayne', 'batsuit.jpg')
print(os.path.split(p))

path = "C:\\Users\\Hyperlight Drifter"

# # only matches on full file name
# filename = "001Bulbasaur.png"
# for dirname, subdirs, files in os.walk(path):
#     if(files.__contains__(filename)):
#         print(f"Path: {dirname}\\{filename}")
#     # print(dirname, 'contains subdirectories:', subdirs, end=' ')
#     # print('and the files:', files)



# read byte stream
path = "C:\\Users\\Hyperlight Drifter\\Documents\\GitHub\\Pokemon-Remaster\\Assets\\Resources\\Images\\PokemonImages\\001Bulbasaur.png"
f = open(path, 'rb')
contents = f.read()
print(contents)
f.close()


file = "ball.bmp"
ball_file = open(file, 'rb')
ball_data = ball_file.read()
ball_file.close()

# BMP image file format stores location
# of pixel RGB values in bytes 10-14
pixel_data_loc = ball_data[10:14]

# Converts byte sequence into integer object
pixel_data_loc = struct.unpack('<L', pixel_data_loc)[0]

# Create sequence of 3000 red, green, and yellow pixels each
new_pixels = b'\x01'*3000 + b'\x02'*3000 + b'\x03'*3000

# Overwrite pixels in image with new pixels
new_ball_data = ball_data[:pixel_data_loc] + \
              new_pixels + \
              ball_data[pixel_data_loc + len(new_pixels):]

# Write new image
new_ball_file = open('new_bmp.bmp', 'wb')
new_ball_file.write(new_ball_data)
new_ball_file.close()

# reverse of pack, returns tuple
print('Result of unpacking {}:'.format(repr('\x00\x05'), end=' '))
print(struct.unpack('>h', b'\x00\x05'))


print('Result of unpacking {}:'.format(repr('\x01\x00'), end=' '))
print(struct.unpack('>h', b'\x01\x00'))

print('Result of unpacking {}:'.format(repr('\x00\x05\x01\x00'), end=' '))
print(struct.unpack('>hh', b'\x00\x05\x01\x00'))


# CSV reading by row
with open('test.csv', 'r') as csvfile:
    csv_reader = csv.reader(csvfile, delimiter=',')

    row_num = 1
    for row in csv_reader:
        print('Row #{}:'.format(row_num), row)
        row_num += 1



# writing rows to a csv
row1 = ['100', '50', '29']
row2 = ['76', '32', '330']

with open('gradeswr.csv', 'w') as csvfile:
    grades_writer = csv.writer(csvfile)

    grades_writer.writerow(row1)
    grades_writer.writerow(row2)

    grades_writer.writerows([row1, row2])