import os
import struct
import pandas as pd
import struct, csv, pickle


# TODO add Pickle

def file_handler(filepath: str, mode: str, contents=None, toDict: bool = False, toCSV: bool = False) -> str:
    file_name = filepath.split(os.path.sep)[-1]
    # print(f'{file_name} size: {os.path.getsize(filepath)} bytes ')
    """
    modes: (r) read (w) write (a) append (b) binary
    (+) read + write
    """
    # with: opens file, execute block and automatically close
    with open(filepath, mode) as f:
        if mode in ['a', 'a+', 'r+', 'w', 'wb', 'w+', 'x']:
            print(f"writing in mode {mode}")
            if contents is not None:
                if file_name.split(".")[-1] == "csv":
                    pass
                    # todo fix
                    # csv_writer = csv.writer(f, delimiter=',', quotechar='"', quoting=csv.QUOTE_ALL)
                    # for i in range(len(contents)):
                    #     print(contents[i])
                    #     csv_writer.writerow(contents[i])
                else:
                    f.write(contents)
        if mode in ['a+', 'r', 'rb', 'r+', 'w+']:
            # print(f"reading in mode {mode}")
            f.seek(0)  # move cursor to the start of the file

            if toDict:
                # TODO return dict if true, example below
                pass
            else:
                return f.read()

        # file_csv = pd.read_csv(pokemon_csv, index_col=0, sep=",", encoding='cp1252')
        # return file_csv.transpose().to_dict(orient='series')


def clean_file_path(filepath: str, cwd: bool = False) -> str:
    # TODO improve with https://stackoverflow.com/questions/13939120/sanitizing-a-file-path-in-python
    clean_path = filepath.replace("\\", os.path.sep).replace("/", os.path.sep)
    if cwd:
        clean_path = os.path.join(os.getcwd(), clean_path)
    return clean_path


def find_file(filename: str, base_dir: str = None) -> str:
    path = base_dir
    if base_dir is None:
        path = os.path.abspath(os.sep)

    # only matches on full file name
    for dirname, subdirs, files in os.walk(path):
        if files.__contains__(filename):
            print(f"Path: {dirname}\\{filename}")
        # print(dirname, 'contains subdirectories:', subdirs, end=' ')
        # print('and the files:', files)


if __name__ == "__main__":

    # Open a file regardless of OS
    # grabs route correctly independent of OS routing:
    #   Linux/Mac: CSV/Pokemon.csv
    #   Windows: CSV\Pokemon.csv
    cwd = os.getcwd()
    pokemon_csv = os.path.join(cwd, 'CSV', "Pokemon.csv")
    base_stats_csv = os.path.join(cwd, 'CSV', "Base_Stats.csv")
    test_txt = os.path.join(cwd, 'CSV', "test.txt")
    test_csv = os.path.join(cwd, 'CSV', "test.csv")
    ball_bmp = os.path.join(cwd, "ball.bmp")

    """file paths"""
    # print(pokemon_csv)
    # print(base_stats_csv)

    """read csv"""
    # print(file_handler(pokemon_csv, 'r'))

    """Example of cleaning file path"""
    # print(clean_file_path("CSV/test.csv", True))
    # print(clean_file_path("C:/Users/Hyperlight Drifter"))

    """overwrite to text.txt"""
    # file_handler(test_txt, "w", "Overwriting Contents")

    """append then read from beginning of file"""
    # file_handler(test_txt, "a+", "\nAdd me some letters!")

    # print(file_handler(test_txt, "r"))

    """file all files with exact match"""
    # filename = "001Bulbasaur.png"
    # find_file(filename)

    """Read byte stream"""
    # print(file_handler(test_csv, 'rb'))

    """using Pandas Library to read csv"""
    # items = pd.read_csv(test_csv).to_dict("records")
    # print(items)

    """Path Exists"""
    # p = os.path.join(os.path.abspath(os.sep), 'Users', 'Hyperlight Drifter', "Documents", "Github")
    # print(f"PATH {p} exists: {os.path.exists(p)}")

    # write output buffer to disk every 100 bytes
    f = open(test_txt, 'w', buffering=100)
    # force output buffer to be written to disk
    f.flush()
    # ensure that all internal buffers associated with f are written to disk
    os.fsync(f.fileno())

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
    new_pixels = b'\x01' * 3000 + b'\x02' * 3000 + b'\x03' * 3000

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
    with open(test_csv, 'r') as csvfile:
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

    # while col_name != '':
    #     with open(csv_data, 'w') as csv_file:
    #         csv_writer = csv.writer(csv_file, delimiter=',', quotechar='"', quoting=csv.QUOTE_ALL)
    #         csv_writer.writerow(col_name)
    #         response = requests.get(url_data).json()
    #         try:
    #             n_response = len(response['states'])
    #         except Exception:
    #             pass
    #         else:
    #             for i in range(n_response):
    #                 info = response['states'][i]
    #                 csv_writer.writerow(info)
    #     time.sleep(sleep_time)
    #     print('Get', len(response['states']), 'data')
