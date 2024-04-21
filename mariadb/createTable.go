package mariadb

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// 根据自定义结构体中定义的生成CREATE TABLE语句
func GenerateCreateTableSQL(data any) (string, error) {
	typ := reflect.TypeOf(data)

	if typ.Kind() != reflect.Struct {
		return "", errors.New("data must be a struct")
	}
	tabName := typ.Name()
	if LowercaseTableName {
		tabName = strings.ToLower(tabName)
	}

	var sql strings.Builder
	sql.WriteString("CREATE TABLE " + tabName + " (\n")

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tagAxis := field.Tag.Get("axis")
		fn1 := field.Name
		ft1 := field.Type
		tagAxisY := field.Tag.Get("axis_y")
		fn2 := field.Name
		ft2 := field.Type
		if LowercaseFieldname {
			fn1 = strings.ToLower(fn1)
			fn2 = strings.ToLower(fn2)
		}
		var tag string
		var fieldName string
		var fieldType reflect.Type
		if tagAxis != "" && tagAxisY == "" {
			tag = tagAxis
			fieldName = fn1
			fieldType = ft1
		} else if tagAxisY != "" && tagAxis == "" {
			tag = tagAxisY
			fieldName = fn2
			fieldType = ft2
		} else {
			continue
		}

		if tag != "" {
			switch fieldType.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				sql.WriteString(fmt.Sprintf("%s INT NOT NULL,\n", fieldName))
			case reflect.String:
				sql.WriteString(fmt.Sprintf("%s VARCHAR(255) NOT NULL,\n", fieldName))
			}
		}
	}

	sql.WriteString("PRIMARY KEY (your_table_primary_key)\n")
	sql.WriteString(");\n")

	return sql.String(), nil
}
