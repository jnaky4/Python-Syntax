package httpTest

import (
	"Go/models"
	"encoding/json"
	"io"
	"net/http"
)

func FetchPokemon(URL string) []models.PokemonTest {
	var pokemon []models.PokemonTest
	get, err := http.Get(URL)
	if err != nil {
		return nil
	}

	defer get.Body.Close()

	all, err := io.ReadAll(get.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(all, &pokemon)
	if err != nil {
		return nil
	}
	return pokemon
}
