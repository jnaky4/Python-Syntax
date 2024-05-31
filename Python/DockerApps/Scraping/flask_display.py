from flask import Flask, render_template
import upcomingGames as s

app = Flask(__name__)


@app.route('/')
def display_game_releases():
    old_releases = s.read_from_file()
    new_releases = s.scrape_gameinformer()
    updated_releases = s.update_game_releases(old_releases, new_releases)
    s.write_to_file(updated_releases)
    return render_template('game_releases.html', game_releases=old_releases)


if __name__ == '__main__':
    app.run(debug=False)
