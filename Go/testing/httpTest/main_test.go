package httpTest

import (
	"Go/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var expected = []models.PokemonTest{
	{
		Dexnum: 1,
		Name:   "Bulbasaur",
		Level:  5,
		Type1:  models.Grass,
		Type2:  models.Poison,
	},
}

var test []struct{

}

var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	marshal, _ := json.Marshal(expected)

	w.Write(marshal)
}))

func TestFetchPermissions(t *testing.T) {
	actual := FetchPokemon(server.URL)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected permissions don't match:\n%+v\n%+v\n", actual, expected)
	}
}