import requests
from bs4 import BeautifulSoup
import urllib3

urllib3.disable_warnings()

def scrape_gameinformer():
    base_url = 'https://www.gameinformer.com'
    routes = ['2024', '2023', '2022', '2021', '2020']

    game_releases = {}


    for r in routes:
        url = f'{base_url}/{r}'
        response = requests.get(url, verify=False)

        if response.status_code == 200:
            gr = {}
            soup = BeautifulSoup(response.text, 'html.parser')
            calendar_entries = soup.find_all('span', class_='calendar_entry')

            for entry in calendar_entries:
                game_name_tag = entry.find('a')
                release_date_tag = entry.find('time')

                # Check if both tags are found
                if game_name_tag and release_date_tag:
                    game_name = game_name_tag.text.strip()
                    release_date = release_date_tag.get('datetime')
                    gr[game_name] = release_date
            game_releases.update(gr)
        else:
            print("Failed to retrieve data.")

    return game_releases


def write_to_file(game_releases):
    with open('game_releases.txt', 'w') as file:
        for game_name, release_date in game_releases.items():
            file.write(f"Game: {game_name}, Release Date: {release_date}\n")

    print("Game releases saved to 'game_releases.txt'")


def read_from_file():
    games_map = {}
    try:
        with open('game_releases.txt', 'r') as file:
            for line in file:
                # Split the line based on the first occurrence of ", "
                split_index = line.strip().find(', ')
                if split_index != -1:
                    game_name = line[:split_index].strip().replace("Game: ", "")
                    release_date = line[split_index + len(', '):].strip().replace("Release Date: ", "")
                    games_map[game_name] = release_date
    except FileNotFoundError:
        print("No game releases found.")
    return games_map


def update_game_releases(old, new):
    # Update the old releases with the new releases
    old.update(new)
    return old


if __name__ == "__main__":
    old_releases = read_from_file()
    new_releases = scrape_gameinformer()

    updated_releases = update_game_releases(old_releases, new_releases)
    write_to_file(updated_releases)

