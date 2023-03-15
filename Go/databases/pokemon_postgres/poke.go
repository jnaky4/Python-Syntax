package pokemon_postgres

import (
	t "Go/time_completion"
	"database/sql"
	"strconv"
)

type Pokemon struct {
	//gorm.Model
	dexnum int
	name string
	type1 string
	type2 string
	stage string
	evolve_level int
	gender_ratio string
	height float32
	weight float32
	description string
	category string
	lvl_speed float32
	base_exp int
	catch_rate int
}
//func (Pokemon) TableName() string {
//	return "Pokedex"
//}

//func GormAllPokemon(db *gorm.DB)(string, error){
//	defer t.Timer()()
//	var poke = Pokemon{}
//	allPokemon := db.Find(&poke)
//	if allPokemon.Error != nil {
//		return "", allPokemon.Error
//	}
//	println(allPokemon.RowsAffected)
//
//	return fmt.Sprintf("%+v\n", allPokemon), nil
//}

//func GormAPokemon(dexnum int, db *gorm.DB)(string, error){
//	defer t.Timer()()
//	var poke = Pokemon{dexnum: dexnum}
//
//	resultPoke := db.First(&poke)
//	if resultPoke.Error != nil {
//		return "", resultPoke.Error
//	}
//
//
//	return fmt.Sprintf("%+v\n", poke), nil
//}

func GetAllPokemon(db *sql.DB) ([]Pokemon, error) {
	defer t.Timer()()
	rows, err := db.Query(`SELECT * FROM "Pokedex";`)
	defer rows.Close()
	if err != nil {
		println("Query Fail")
		return nil, err
	}
	var poke Pokemon
	var allpoke []Pokemon

	for rows.Next() {
		//err := rows.Scan(&poke) splat operator?
		err := rows.Scan(&poke.dexnum, &poke.name, &poke.type1, &poke.type2, &poke.stage, &poke.evolve_level, &poke.gender_ratio, &poke.height, &poke.weight, &poke.description, &poke.category, &poke.lvl_speed, &poke.base_exp, &poke.catch_rate)
		if err != nil {
			println("Scan fail")
			return nil, err
		}
		allpoke = append(allpoke, poke)
	}
	return allpoke, nil
}

func GetAPokemon(dexnum string, db *sql.DB) (*Pokemon, error) {
	defer t.Timer()()
	atoi, err := strconv.Atoi(dexnum)
	if err != nil || atoi < 1{
		println("Invalid Dexnum")
		return nil, err
	}
	var poke Pokemon
	query := `SELECT * FROM "Pokedex" WHERE dexnum=` + dexnum
	rows := db.QueryRow(query)

	err = rows.Scan(&poke.dexnum, &poke.name, &poke.type1, &poke.type2, &poke.stage, &poke.evolve_level, &poke.gender_ratio, &poke.height, &poke.weight, &poke.description, &poke.category, &poke.lvl_speed, &poke.base_exp, &poke.catch_rate)
	if err != nil {
		println("Scan fail")
		return nil, err
	}
	return &poke, nil
}