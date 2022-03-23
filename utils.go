package main

import (
	"bytes"
	"reflect"
	"regexp"
	"strings"
)

func toSnake(s string) string {
	buf := bytes.NewBufferString("")
	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' {
			buf.WriteRune('_')
		}
		buf.WriteRune(v)
	}
	return strings.ToLower(buf.String())
}

func ModelExtract(model interface{}) (tableName string, columns []string, values []interface{}) {
	typ := reflect.TypeOf(model).Elem()
	tableName = toSnake(typ.Name())
	reg, _ := regexp.Compile("s*$")
	tableName = reg.ReplaceAllString(tableName, "s")

	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			columns = append(columns, p.Name)
			value := reflect.ValueOf(model).Elem().FieldByName(p.Name)
			values = append(values, value.Interface())
		}
	}
	return
}
