package mariadb

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func getUpdateFields(tableName string, data any, fieldCustomTag string) (string, error) {
	typ := reflect.TypeOf(data)

	if typ.Kind() != reflect.Struct {
		return "", errors.New("data must be a struct")
	}
	var tabName string
	if tableName == "" {
		tabName = typ.Name() // 获取结构体的名称作为表名
	} else {
		tabName = tableName
	}
	if LowercaseTableName {
		tabName = strings.ToLower(tabName)
	}

	var ins strings.Builder
	ins.WriteString("INSERT INTO " + tabName + " (")
	var sql strings.Builder
	sql.WriteString("UPDATE " + tabName + " SET ")
	fieldNames := []string{}
	fieldNamesOnly := []string{}
	for i := 0; i < typ.NumField(); i++ {
		var tag string
		var fieldName string
		field := typ.Field(i)
		switch fieldCustomTag {
		case "":
			tagAxis := field.Tag.Get("axis")
			fn1 := field.Name
			tagAxisY := field.Tag.Get("axis_y")
			fn2 := field.Name
			if LowercaseFieldname {
				fn1 = strings.ToLower(fn1)
				fn2 = strings.ToLower(fn2)
			}
			if tagAxis != "" && tagAxisY == "" {
				tag = tagAxis
				fieldName = fn1
			} else if tagAxisY != "" && tagAxis == "" {
				tag = tagAxisY
				fieldName = fn2
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
		default:
			tag = field.Tag.Get(fieldCustomTag)
			if LowercaseFieldname {
				fieldName = strings.ToLower(tag)
			} else {
				fieldName = tag
			}
		}

		if tag != "" {
			fieldNamesOnly = append(fieldNamesOnly, fieldName)
			fieldNames = append(fieldNames, fmt.Sprintf("%s=?", fieldName))
		}
	}
	ins.WriteString(strings.Join(fieldNamesOnly, ", "))
	ins.WriteString(") VALUES (")
	_insSplit := make([]string, len(fieldNamesOnly))
	for i := 0; i < len(fieldNames); i++ {
		_insSplit[i] = "?"
	}
	ins.Write([]byte(strings.Join(_insSplit, ", ")))
	ins.WriteByte(')')
	ins.WriteByte('\n')
	ins.WriteByte('\n')

	sql.WriteString(strings.Join(fieldNames, ", "))
	sql.WriteByte('\n')
	return ins.String() + sql.String(), nil
}

func GetUpdateFields(data any) (string, error) {
	return getUpdateFields("", data, "")
}

func GetUpdateAllField(data any) (string, error) {
	return getUpdateFields("", data, "*")
}

func GetUpdateFieldsFromCustomTag(data any, fieldTag string) (string, error) {
	if fieldTag == "" || fieldTag == "*" {
		return "", errors.New("fieldTag must not be empty")
	}
	return getUpdateFields("", data, fieldTag)
}

func GetUpdateFieldsNeedTabName(tableName string, data any) (string, error) {
	if tableName == "" {
		return "", errors.New("tableName must not be empty")
	}
	return getUpdateFields(tableName, data, "")
}

func GetUpdateAllFieldNeedTabName(tableName string, data any) (string, error) {
	if tableName == "" {
		return "", errors.New("tableName must not be empty")
	}
	return getUpdateFields(tableName, data, "*")
}

func GetUpdateFieldsFromCustomTagNeedTabName(tableName string, data any, fieldTag string) (string, error) {
	if tableName == "" {
		return "", errors.New("tableName must not be empty")
	}
	if fieldTag == "" || fieldTag == "*" {
		return "", errors.New("fieldTag must not be empty")
	}
	return getUpdateFields(tableName, data, fieldTag)
}
