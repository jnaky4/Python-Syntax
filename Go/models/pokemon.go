package models

import (
	"strings"
)

//type Type uint8
//const (
//	Bug Type = iota
//	Dragon
//	Electric
//	Fighting
//	Fire
//	Flying
//	Ghost
//	Grass
//	Ground
//	Ice
//	Normal
//	Poison
//	Psychic
//	Rock
//	Water
//)
//func (t Type) String() string {
//	return [...]string{"Bug", "Dragon", "Electric", "Fighting", "Fire", "Flying", "Ghost", "Gras", "Ground", "Ice", "Normal", "Poison", "Psychic", "Rock", "Water"}[t-1]
//}

type Type string

const (
	Fire     Type = "fire"
	Water    Type = "water"
	Grass    Type = "grass"
	Electric Type = "electric"
	Ice      Type = "ice"
	Fighting Type = "fighting"
	Poison   Type = "poison"
	Ground   Type = "ground"
	Flying   Type = "flying"
	Psychic  Type = "psychic"
	Bug      Type = "bug"
	Rock     Type = "rock"
	Ghost    Type = "ghost"
	Dragon   Type = "dragon"
	Normal   Type = "normal"
)

var typeChart = map[Type]map[Type]float64{
	Fire: {
		Fire: .5,
		Water: .5,
		Grass: 2,
		Ice: 2,
		Bug: 2,
		Rock: .5,
		Dragon: .5,
	},
	Water: {
		Fire: 2,
		Water: .5,
		Grass: .5,
		Ground: 2,
		Rock: 2,
		Dragon: .5,
	},
	Grass: {
		Fire: .5,
		Water: 2,
		Grass: .5,
		Poison: .5,
		Ground: 2,
		Flying: .5,
		Bug: .5,
		Rock: 2,
		Dragon: .5,
	},
	Electric: {
		Water: 2,
		Grass: .5,
		Electric: .5,
		Ground: 0,
		Flying: 2,
		Dragon: .5,
	},
	Ice: {
		Water: .5,
		Grass: 2,
		Ice: .5,
		Ground: 2,
		Flying: 2,
		Dragon: 2,
	},
	Fighting: {
		Ice: 2,
		Poison: .5,
		Flying: .5,
		Psychic: .5,
		Bug: .5,
		Rock: 2,
		Ghost: 0,
		Normal: 2,
	},
	Poison: {
		Grass: 2,
		Poison: .5,
		Ground: .5,
		Bug: 2,
		Rock: .5,
		Ghost: .5,
	},
	Ground: {
		Fire: 2,
		Grass: .5,
		Electric: 2,
		Poison: 2,
		Flying: 0,
		Bug: .5,
		Rock: 2,
	},
	Flying: {
		Grass: 2,
		Electric: .5,
		Fighting: 2,
		Bug: 2,
		Rock: .5,
	},
	Psychic: {
		Fighting: 2,
		Poison: 2,
		Psychic: .5,
	},
	Bug: {
		Fire: .5,
		Grass: 2,
		Fighting: .5,
		Poison: 2,
		Flying: .5,
		Psychic: 2,
		Ghost: .5,
	},
	Rock:{
		Fire: 2,
		Ice: 2,
		Fighting: .5,
		Ground: .5,
		Flying: 2,
		Bug: 2,
	},
	Ghost: {
		Psychic: 2,
		Ghost: 2,
		Normal: 0,
	},
	Dragon: {
		Dragon: 2,
	},
	Normal: {
		Rock: .5,
		Ghost: 0,
	},
}

type Effectiveness struct{
	Immune []string `json:"imm,omitempty" csv:"imm" sql:"imm"`
	SResist []string `json:"sresist,omitempty" csv:"sresist" sql:"sresist"`
	Resist []string `json:"resist,omitempty" csv:"resist" sql:"resist"`
	Weak []string	`json:"weak,omitempty" csv:"weak" sql:"weak"`
	SWeak []string	`json:"sweak,omitempty" csv:"sweak" sql:"sweak"`
}

type BaseStats struct{
	Id int `json:"dexnum" csv:"dexnum" sql:"dexnum"`
	Hp int `json:"hp" csv:"hp" sql:"hp"`
	Atk int `json:"atk" csv:"atk" sql:"atk"`
	SpAtk int `json:"satk" csv:"satk" sql:"satk"`
	Def int `json:"def" csv:"def" sql:"def"`
	SpDef int `json:"sdef" csv:"sdef" sql:"sdef"`
	Spd int `json:"spd" csv:"spd" sql:"spd"`
	Total int `json:"total" csv:"total" sql:"total"`
}

type Pokedex struct{
	Stage       string  `json:"stage" csv:"stage" sql:"stage"`
	EvolveLevel int     `json:"evolve_level" csv:"evolve_level" sql:"evolve_level"`
	GenderRatio string  `json:"gender_ratio" csv:"gender_ratio" sql:"gender_ratio"`
	Height      float32 `json:"height" csv:"height" sql:"height"`
	Weight      float32 `json:"weight" csv:"weight" sql:"weight"`
	Description string  `json:"description" csv:"description" sql:"description"`
	Category    string  `json:"category" csv:"category" sql:"category"`
	LvlSpeed    float32 `json:"lvl_speed" csv:"lvl_speed" sql:"lvl_speed"`
	BaseExp     int     `json:"base_exp" csv:"base_exp" sql:"base_exp"`
	CatchRate   int     `json:"catch_rate" csv:"catch_rate" sql:"catch_rate"`
}

type Pokemon struct {
	Dexnum      int     `json:"dexnum" csv:"dexnum" sql:"dexnum"`
	Name        string  `json:"name" csv:"name" sql:"name"`
	Type1       string  `json:"type1" csv:"type1" sql:"type1"`
	Type2       string  `json:"type2,omitempty" csv:"type2" sql:"type2"`
	Pokedex
	Effectiveness
	BaseStats
}

type PokemonTest struct {
	Dexnum int
	Name   string
	Type1  Type
	Type2  Type
	Level  int
}

func (p* Pokemon) CalculateTypeEffectiveness() {

	for k := range typeChart{
		var score float64
		if p.Type2 == ""{
			score = TypeResult(string(k), strings.ToLower(p.Type1))
		} else{
			score = TypeResult(string(k), strings.ToLower(p.Type1)) * TypeResult(string(k), strings.ToLower(p.Type2))
		}

		switch score {
		case 0:
			p.Immune = append(p.Immune, string(k))
		case .25:
			p.SResist = append(p.SResist, string(k))
		case .5:
			p.Resist = append(p.Resist, string(k))
		case 2:
			p.Weak = append(p.Weak, string(k))
		case 4:
			p.SWeak = append(p.SWeak, string(k))
		}
	}
}

func TypeResult(attackType string, defendType string) float64{
	if _, exists := typeChart[Type(attackType)][Type(defendType)]; !exists {
		return 1
	}
	return typeChart[Type(attackType)][Type(defendType)]
}

func MergePokemonStructs(poke []Pokemon, stats []BaseStats)[]Pokemon{
	for i := range poke{
		if poke[i].Dexnum == stats[i].Id {
			poke[i].BaseStats = stats[i]
		}
	}
	return poke
}