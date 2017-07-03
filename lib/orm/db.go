package orm

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"
)

// All select rows
func All(db *sql.DB, hnd func(*sql.Rows) error, query string, args ...interface{}) error {
	log.Info(query)
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := hnd(rows); err != nil {
			return err
		}
	}
	return rows.Err()
}

// One select row
func One(db *sql.DB, hnd func(*sql.Row) error, query string, args ...interface{}) error {
	log.Info(query)
	row := db.QueryRow(query, args...)
	return hnd(row)
}

// Do execute
func Do(db *sql.DB, query string, args ...interface{}) (int64, int64, error) {
	log.Info(query)
	rst, err := db.Exec(query, args...)
	if err != nil {
		return 0, 0, err
	}
	insertID, err := rst.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	affected, err := rst.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	return insertID, affected, err
}
