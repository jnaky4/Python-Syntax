package pg_util

import (
	"database/sql"
	"fmt"
)

type pg_config struct {
	op command
	table string
	DB *sql.DB
	args []interface{}
}
type command string

const (
	ALTER command = "ALTER"
	CREATE command = "CREATE"
	DELETE command = "DELETE"
	INSERT command = "INSERT"
	SELECT command = "SELECT"
	UPDATE command = "UPDATE"
)

func (p *pg_config)Insert() error {

	query := `INSERT INTO warmState(image, stores)
	VALUES ($1, $2)
	RETURNING id, created_at, version`

	p.args = []interface{}{}

	return p.DB.QueryRow(query, p.args...).Scan()
	return nil
}

func (p *pg_config)Query() error {
	query := ""
	p.args = []interface{}{}

	switch p.op {
		case INSERT:
			query = fmt.Sprintf("%s INTO %s()", p.op, p.table)

	}


	//query := `INSERT INTO warmState(image, stores)
	//VALUES ($1, $2)
	//RETURNING id, created_at, version`
	//
	//args := []interface{}{}
	return p.DB.QueryRow(query, p.args...).Scan()
	return nil

}