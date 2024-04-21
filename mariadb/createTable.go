package mariadb

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// 根据自定义结构体中定义的生成CREATE TABLE语句
func generateCreateTableSQL(data any, fieldCustomTag string) (string, error) {
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
		var tag string
		var fieldName string
		var fieldType reflect.Type
		field := typ.Field(i)
		switch fieldCustomTag {
		case "":
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
		case "*":
			tag = "*"
			if LowercaseFieldname {
				fieldName = strings.ToLower(field.Name)
			} else {
				fieldName = field.Name
			}
			fieldType = field.Type
		default:
			tag = field.Tag.Get(fieldCustomTag)
			if LowercaseFieldname {
				fieldName = strings.ToLower(tag)
			} else {
				fieldName = tag
			}
			fieldType = field.Type
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

// 根据自定义结构体中定义的Tag默认值生成CREATE TABLE语句。tag=="axis" 或 tag=="axis_y"
func GenerateCreateTableSQL(data any) (string, error) {
	return generateCreateTableSQL(data, "")
}

// 根据自定义结构的所有字段生成CREATE TABLE语句。
func GenerateCreateTableSQLAllField(data any) (string, error) {
	return generateCreateTableSQL(data, "*")
}

// 根据自定义结构体,指定Tag属性名 生成CREATE TABLE语句。
func GenerateCreateTableSQLCustomTag(data any, fieldTag string) (string, error) {
	return generateCreateTableSQL(data, fieldTag)
}
