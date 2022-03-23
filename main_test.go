package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelExtract(t *testing.T) {
	var tmp TTT
	tmp.AA = 1
	a, b, c := ModelExtract(&tmp)
	fmt.Println(a, b, c)
}

func TestOpen(t *testing.T) {
	db, err := Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/tuya_app_uitest?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8")
	assert.Nil(t, err)
	err = db.Ping()
	assert.Nil(t, err)
}
