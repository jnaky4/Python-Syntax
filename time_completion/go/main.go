package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func scrapeGameInformer() map[string]string {
	baseURL := "https://www.gameinformer.com"
	routes := []string{"2024", "2023", "2022", "2021", "2020"}

	gameReleases := make(map[string]string)

	for _, r := range routes {
		url := fmt.Sprintf("%s/%s", baseURL, r)
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to retrieve data for route %s\n", r)
			continue
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			doc, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				fmt.Printf("Failed to parse HTML for route %s\n", r)
				continue
			}

			doc.Find("span.calendar_entry").Each(func(i int, s *goquery.Selection) {
				gameName := strings.TrimSpace(s.Find("a").Text())
				releaseDate, _ := s.Find("time").Attr("datetime")
				gameReleases[gameName] = releaseDate
			})
		} else {
			fmt.Printf("Failed to retrieve data for route %s. Status code: %d\n", r, response.StatusCode)
		}
	}

	return gameReleases
}

func writeToFile(gameReleases map[string]string) {
	file, err := os.Create("game_releases.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	for gameName, releaseDate := range gameReleases {
		_, err := fmt.Fprintf(file, "Game: %s, Release Date: %s\n", gameName, releaseDate)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}
	}

	fmt.Println("Game releases saved to 'game_releases.txt'")
}

func readFromFile() map[string]string {
	gameReleases := make(map[string]string)

	file, err := os.Open("game_releases.txt")
	if err != nil {
		fmt.Println("No game releases found.")
		return gameReleases
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read from file:", err)
		return gameReleases
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ", ")
		if len(fields) == 2 {
			gameName := strings.TrimSpace(strings.TrimPrefix(fields[0], "Game: "))
			releaseDate := strings.TrimSpace(strings.TrimPrefix(fields[1], "Release Date: "))
			gameReleases[gameName] = releaseDate
		}
	}

	return gameReleases
}

func updateGameReleases(old, new map[string]string) map[string]string {
	for k, v := range new {
		old[k] = v
	}
	return old
}

func main() {
	//oldReleases := readFromFile()
	newReleases := scrapeGameInformer()

	//updatedReleases := updateGameReleases(oldReleases, newReleases)
	writeToFile(newReleases)
}
