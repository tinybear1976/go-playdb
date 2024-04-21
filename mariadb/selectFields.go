package mariadb

import (
	"errors"
	"reflect"
	"strings"
)

func getFields(data any, fieldCustomTag string) (string, error) {
	typ := reflect.TypeOf(data)

	if typ.Kind() != reflect.Struct {
		return "", errors.New("data must be a struct")
	}
	tabName := typ.Name()
	if LowercaseTableName {
		tabName = strings.ToLower(tabName)
	}

	var sql strings.Builder
	sql.WriteString("SELECT ")
	fieldNames := []string{}
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
			fieldNames = append(fieldNames, fieldName)
		}
	}
	sql.WriteString(strings.Join(fieldNames, ", "))
	sql.WriteString("\nFROM ")
	sql.WriteString(tabName)

	return sql.String(), nil
}

func GetSelectFields(data any) (string, error) {
	return getFields(data, "")
}

func GetSelectAllField(data any) (string, error) {
	return getFields(data, "*")
}

func GetSelectFieldsFromCustomTag(data any, fieldTag string) (string, error) {
	if fieldTag == "" || fieldTag == "*" {
		return "", errors.New("fieldTag must not be empty")
	}
	return getFields(data, fieldTag)
}
