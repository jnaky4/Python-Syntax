package pokemon_postgres

import (
	m "Go/models"
	t "Go/time_completion"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-pg/pg/v10/orm"
	//sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	//"os"
)

type pgModel struct {
	DB *sql.DB
}

type Models struct {
	Pokemon pgModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Pokemon: pgModel{DB: db},
	}
}

func (p pgModel) Insert(poke m.Pokemon) error {
	//defer t.FunctionTimer(p.Insert)()
	query := `
	INSERT INTO pokemon (
	                     dexnum, 
	                     name, 
	                     type1, 
	                     type2, 
	                     stage, 
	                     evolve_level, 
	                     gender_ratio, 
	                     height, 
	                     weight, 
	                     description, 
	                     category, 
	                     lvl_speed, 
	                     base_exp, 
	                     catch_rate
	                     )
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	args := []interface{}{poke.Dexnum, poke.Name, poke.Type1, poke.Type2, poke.Stage, poke.EvolveLevel, poke.GenderRatio,
		poke.Height, poke.Weight, poke.Description, poke.Category, poke.LvlSpeed, poke.BaseExp, poke.CatchRate}
	_, err := p.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil

}

func (p pgModel) Get(dexnum int, count int) (allPoke []m.Pokemon, err error) {
	defer t.FunctionTimer(p.Get)()
	var query string
	var rows *sql.Rows
	tn := "pokemon"

	if dexnum == 0 {
		query = fmt.Sprintf(`SELECT * FROM %s `, tn)
	} else if count == 0 {
		query = fmt.Sprintf(`SELECT * FROM %s WHERE dexnum=%d`, tn, dexnum)
	} else {
		query = fmt.Sprintf(`SELECT * FROM %s WHERE dexnum BETWEEN %d AND %d`, tn, dexnum, count)
	}

	rows, err = p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var poke m.Pokemon
	for rows.Next() {
		//err := rows.Scan(&poke) splat operator?
		err = rows.Scan(
			&poke.Dexnum,
			&poke.Name,
			&poke.Type1,
			&poke.Type2,
			&poke.Stage,
			&poke.EvolveLevel,
			&poke.GenderRatio,
			&poke.Height,
			&poke.Weight,
			&poke.Description,
			&poke.Category,
			&poke.LvlSpeed,
			&poke.BaseExp,
			&poke.CatchRate,
		)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, errors.New("record not found")
			default:
				return nil, err
			}

		}
		allPoke = append(allPoke, poke)
	}
	return allPoke, nil
}

func (p pgModel) Update(poke m.Pokemon) error {
	query := `
		UPDATE Pokemon
		SET Id = $1,
		Name = $2,
		Type1 = $3,
		Type2 = $4,
		Stage = $5,
		EvolveLevel = $6,
		GenderRatio = $7,
		Height = $8,
		Weight = $9,
		Description = $10,
		Category = $11,
		LvlSpeed = $12,
		BaseExp = $13,
		CatchRate = $14,
		WHERE dexnum = $1
	`
	args := []interface{}{
		poke.Dexnum,
		poke.Name,
		poke.Type1,
		poke.Type2,
		poke.Stage,
		poke.EvolveLevel,
		poke.GenderRatio,
		poke.Height,
		poke.Weight,
		poke.Description,
		poke.Category,
		poke.LvlSpeed,
		poke.BaseExp,
		poke.CatchRate,
	}
	return p.DB.QueryRow(query, args...).Scan()
}

func (p pgModel) Delete(dexnum int) error {
	query := `
	DELETE FROM Pokemon
	Where Id = $1
	`
	r, err := p.DB.Exec(query, dexnum)
	if err != nil {
		return err
	}
	affected, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("record not found")
	}
	return nil
}

func (p pgModel) Len(tn string) (size int, err error) {
	err = p.DB.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM ` + tn)).Scan(&size)
	return
}

func (p pgModel) CreatePokemonTable() error {
	tn := "pokemon"
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ("+
		"dexnum INT NOT NULL PRIMARY KEY, "+
		"name TEXT NOT NULL,"+
		"type1 TEXT NOT NULL,"+
		"type2 TEXT,"+
		"stage TEXT NOT NULL,"+
		"evolve_level INT,"+
		"gender_ratio TEXT NOT NULL,"+
		"height FLOAT NOT NULL,"+
		"weight FLOAT NOT NULL,"+
		"description TEXT NOT NULL,"+
		"category TEXT NOT NULL,"+
		"lvl_speed TEXT NOT NULL,"+
		"base_exp INT NOT NULL,"+
		"catch_rate INT NOT NULL"+
		")", tn)
	_, err := p.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
