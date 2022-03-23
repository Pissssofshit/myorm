package main

import (
	"database/sql"

	_ "gorm.io/driver/mysql"
)

type DB struct {
	db     *sql.DB
	sql    string
	params []interface{}
}

func NewDB(source string) (*DB, error) {
	db, err := Open("mysql", source)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

func Open(driver string, source string) (*sql.DB, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (o *DB) Save() {
	o.Explain()
	o.db.Exec(o.sql, o.params...)
}

func (o *DB) Explain() {
	// do something to sql & params
}
