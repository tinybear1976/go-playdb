package goplaydb

import (
	"errors"
	"reflect"
	"strings"
)

func GoObjectReference(structName string, struc any, fieldCustomTag string) (string, error) {
	typ := reflect.TypeOf(struc)

	if typ.Kind() != reflect.Struct {
		return "", errors.New("data must be a struct")
	}
	var rawStructName string
	if structName == "" {
		rawStructName = typ.Name() // 获取结构体的名称作为表名
	} else {
		rawStructName = structName
	}
	var goCodeInsert strings.Builder
	var goCodeUpdate strings.Builder
	/* 类似以下结构
		func (obj Student) InsertToDB() error {
		   obj.IsShow ,
		   return nil
	  }
	*/
	goCodeInsert.WriteString("// golang code \n\n")
	goCodeInsert.WriteString("func (obj " + rawStructName + ") InsertToDB() error {\n")
	goCodeUpdate.WriteString("func (obj " + rawStructName + ") UpdateToDB() error {\n")
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
			fieldName = field.Name
		default:
			tag = field.Tag.Get(fieldCustomTag)
			fieldName = field.Name
		}

		if tag != "" {
			fieldNames = append(fieldNames, "obj."+fieldName)
		}
	}
	goCodeInsert.WriteString(strings.Join(fieldNames, ", "))
	goCodeInsert.WriteByte('\n')
	goCodeUpdate.WriteString(strings.Join(fieldNames, ", "))
	goCodeUpdate.WriteByte('\n')

	goCodeInsert.WriteString("return nil\n}\n")
	goCodeUpdate.WriteString("return nil\n}\n")
	return goCodeInsert.String() + "\n" + goCodeUpdate.String() + "\n", nil
}
