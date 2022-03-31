package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

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

func (o *DB) CreateTable(model interface{}) (err error) {
	_, err = o.db.Exec(CreateSQL(model))
	return
}

func (d *DB) SaveSql(model interface{}) string {
	typ := reflect.TypeOf(model).Elem()
	tableName := strings.ToLower(typ.Name())
	sql := "INSERT TABLE `" + tableName + "` ("

	var columnNames []string
	var values []interface{}
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			colunmName := strings.ToLower(p.Name)
			columnNames = append(columnNames, colunmName)
			values = append(values, reflect.ValueOf(model).Elem().FieldByName(p.Name).Interface())
		}
	}
	for _, col := range columnNames {
		sql += col + ","
	}
	sql = strings.TrimRight(sql, ",")
	sql += ")"

	sql += " values("
	// for _, value := range values {
	// 	sql += value + ","
	// }
	fmt.Println(values...)

	d.params = values
	fmt.Println(sql)
	return sql
}

func CreateSQL(model interface{}) string {
	typ := reflect.TypeOf(model).Elem()
	//TODO
	tableName := strings.ToLower(typ.Name())
	sql := "CREATE TABLE `" + tableName + "` ("
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			colunmName := strings.ToLower(p.Name)
			value := reflect.ValueOf(model).Elem().FieldByName(p.Name).Interface()
			sqlType := getSqlType(value)
			sql += colunmName + " " + sqlType + ","
		}
	}
	sql = strings.TrimRight(sql, ",")
	sql += ")"
	fmt.Println(sql)
	return sql
}

func (o *DB) Save(model interface{}) (err error) {
	_, err = o.db.Exec(SaveSql(model))
	return
}

func (o *DB) Explain(model interface{}) {
	// do something to sql & params
}

func getSqlType(column interface{}) string {
	// 这种写法是switch 专用
	switch column.(type) {
	case time.Time:
		return "timestamp"
	case bool:
		return "boolean"
	case int, int8, int16, int32, uint, uint8, uint16, uint32:
		return "int"
	case int64, uint64:
		return "bigint"
	case float32, float64:
		return "double"
	case string:
		return "varchar(255)"
	default:
		panic("invalid sql type")
	}
}
