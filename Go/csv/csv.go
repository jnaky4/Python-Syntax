package csv

import (
	m "Go/models"
	"github.com/gocarina/gocsv"
	"os"
)

func LoadPokemonCSV(fPath string) ([]m.Pokemon, error) {
	var poke []m.Pokemon
	file, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	err = gocsv.UnmarshalFile(file, &poke)
	if err != nil {
		return nil, err
	}
	return poke, nil
}

func LoadBaseStatsCSV(fPath string) ([]m.BaseStats, error){
	var stats []m.BaseStats
	file, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	err = gocsv.UnmarshalFile(file, &stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}


