package pokemon_postgres

import (
	t "Go/time_completion"
	"database/sql"
	"strconv"
)

type Pokemon struct {
	//gorm.Model
	Dexnum       int
	Name         string
	Type1        string
	Type2        string
	Stage        string
	Evolve_level int
	Gender_ratio string
	Height       float32
	Weight       float32
	Description  string
	Category     string
	Lvl_speed    float32
	Base_exp     int
	Catch_rate   int
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

//func GormAPokemon(Dexnum int, db *gorm.DB)(string, error){
//	defer t.Timer()()
//	var poke = Pokemon{Dexnum: Dexnum}
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
	defer t.FunctionTimer(GetAllPokemon)()
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
		err := rows.Scan(&poke.Dexnum, &poke.Name, &poke.Type1, &poke.Type2, &poke.Stage, &poke.Evolve_level, &poke.Gender_ratio, &poke.Height, &poke.Weight, &poke.Description, &poke.Category, &poke.Lvl_speed, &poke.Base_exp, &poke.Catch_rate)
		if err != nil {
			println("Scan fail")
			return nil, err
		}
		allpoke = append(allpoke, poke)
	}
	return allpoke, nil
}

func GetAPokemon(dexnum string, db *sql.DB) (*Pokemon, error) {
	defer t.FunctionTimer(GetAPokemon)()
	atoi, err := strconv.Atoi(dexnum)
	if err != nil || atoi < 1 {
		println("Invalid Dexnum")
		return nil, err
	}
	var poke Pokemon
	query := `SELECT * FROM "Pokedex" WHERE Dexnum=` + dexnum
	rows := db.QueryRow(query)

	err = rows.Scan(&poke.Dexnum, &poke.Name, &poke.Type1, &poke.Type2, &poke.Stage, &poke.Evolve_level, &poke.Gender_ratio, &poke.Height, &poke.Weight, &poke.Description, &poke.Category, &poke.Lvl_speed, &poke.Base_exp, &poke.Catch_rate)
	if err != nil {
		println("Scan fail")
		return nil, err
	}
	return &poke, nil
}
