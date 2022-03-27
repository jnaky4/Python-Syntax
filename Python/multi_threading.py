from sounds import get_a_pokemon_sound, get_pokemon_sounds
import os
from multiprocessing import Pool

# single threaded network - io bound process
# get_pokemon_sounds()

# for j in range(1, 152):
#     play_pokemon_sound(j)

# multiple threaded network - io bound process
cpus = os.cpu_count()

urls = []
for j in range(1, 152):
    urls.append(f"https://pokemoncries.com/cries-old/{j}.mp3")

with Pool(cpus) as pool:
    res = pool.map(get_a_pokemon_sound, urls)