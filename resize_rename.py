import os
from PIL import Image

dir = "C:/Users/Hyperlight Drifter/Documents/GitHub/Pokemon-Remaster/Assets/Resources/Images/PokemonImages/Rename/"

# files = os.listdir(dir)

# for file in files:
#     if re.search('.+png', file):
#         print(file)

for filename in os.listdir(dir):
    #does path + filename2.png exist dont make it
    if(      (not (os.path.exists(dir+filename[:-4]+'2.png'))) or (not (os.path.exists(dir+filename[:-5]+'2.png')))       ):
        original_img = Image.open(dir+filename)
        original_img.save(dir+filename[:-4]+'2.png')
    # horz_image = original_img.transpose(method=Image.FLIP_LEFT_RIGHT)
    # horz_image.save(dir+filename[:-4]+"2.png")
        # # print(filename, filename[:-6], filename[-4:])
        # os.rename(dir+filename, dir+(filename[:-6]+filename[-4:]))
