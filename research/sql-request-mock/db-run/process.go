package db_run

import (
	"database/sql"
	"errors"
)

const query = `select
	    key_words,
	    views,
	    clicks,
	from
	    table`

type dbConn interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

func Run(db dbConn) error {
	rows, errQuery := db.Query(query)
	if errQuery != nil {
		return errQuery
	}

	for rows.Next() {
		var (
			key_words string
			views     int
			clicks    int
		)

		err := rows.Scan(&key_words, &views, &clicks)
		if err != nil {
			return err
		}

		if key_words == "error_words1" {
			return errors.New("error_words")
		}
	}

	return nil
}
