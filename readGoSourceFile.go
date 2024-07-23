package goplaydb

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/tinybear1976/go-playdb/mariadb"
)

func readGoSourceFile(fileName string) (string, error) {
	// 读取文件内容
	content, err := os.ReadFile(fileName)
	if err != nil {
		//fmt.Println("Error reading file:", err)
		return "", err
	}
	// 输出文件内容
	return string(content), nil
}

type StructInfo struct {
	FieldName string
	FieldType string
	Tag       string
}

func GetStructs(fileName string) (map[string][]StructInfo, error) {
	txt, err := readGoSourceFile(fileName)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", txt, 0)
	if err != nil {
		return nil, err
	}
	// ast.Print(fset, f)

	objects := make(map[string][]StructInfo)
	var tmp []StructInfo = nil
	var tmpName string = ""
	ast.Inspect(f, func(x ast.Node) bool {
		typ, typOk := x.(*ast.TypeSpec)
		if typOk {
			// fmt.Println("type name:", typ.Name.Name)
			objects[typ.Name.Name] = nil
			if tmp != nil {
				des := make([]StructInfo, len(tmp))
				copy(des, tmp)
				objects[tmpName] = des
				tmp = nil
			}
			tmpName = typ.Name.Name
			return true
		}
		s, ok := x.(*ast.StructType)
		if !ok {
			return true
		}
		if tmp == nil {
			tmp = make([]StructInfo, 0)
		}
		for _, field := range s.Fields.List {
			var _tag string
			if field.Tag == nil {
				_tag = ""
			} else {
				_tag = field.Tag.Value
			}
			// fmt.Printf("%#v", field)
			_st := StructInfo{
				FieldName: field.Names[0].Name,
				FieldType: fmt.Sprintf("%v", field.Type),
				Tag:       _tag,
			}
			// fmt.Println(_st)
			tmp = append(tmp, _st)

			// fmt.Printf("Field: %s\n", field.Names[0].Name)
			// fmt.Printf("Tag:   %s\n", field.Tag.Value)
		}
		return false
	})
	if tmp != nil {
		des := make([]StructInfo, len(tmp))
		copy(des, tmp)
		objects[tmpName] = des
	}
	// //fmt.Printf("%#v\n", objects)
	return objects, nil
	// return nil, nil
}

func reBuildStruct(sis []StructInfo) any {
	fields := []reflect.StructField{}
	for _, si := range sis {
		f := reflect.StructField{
			Name: si.FieldName,
			Tag:  reflect.StructTag(si.Tag),
		}
		switch si.FieldType {
		case "int":
			f.Type = reflect.TypeOf(int(0))
		case "string":
			f.Type = reflect.TypeOf("")
		case "float64":
			f.Type = reflect.TypeOf(float64(0))
		case "bool":
			f.Type = reflect.TypeOf(bool(false))
		case "&{decimal Decimal}":
			f.Type = reflect.TypeOf(decimal.NewFromInt(0))
			//	fmt.Println("//////", si.FieldType)
		}
		fields = append(fields, f)
	}
	typ := reflect.StructOf(fields)
	structInstance := reflect.New(typ).Elem()
	return structInstance.Interface()
}

// 依照go结构体文件生成sql文件.
//
//	fileName  为go结构体文件名
//	tagMode   为结构体字段tag模式，默认为空,表示配合 标签"axis" 或 "axis_y"标签生成字段. *生成所有字段，其他为自定义tag
func GenerateSQLFromFile(fileName string, tagMode string) error {
	objs, err := GetStructs(fileName)
	if err != nil {
		return err
	}
	_baseName := filepath.Base(fileName)
	_baseName = strings.TrimSuffix(_baseName, filepath.Ext(_baseName))
	sqlFilename := _baseName + ".sql"
	for tbName, sis := range objs {
		structInstance := reBuildStruct(sis)
		switch tagMode {
		case "":
			err = batchDefault(tbName, sqlFilename, structInstance)
		case "*":
			err = batchAllFields(tbName, sqlFilename, structInstance)
		default:
			err = batchCustomTag(tbName, sqlFilename, structInstance, tagMode)
		}
		if err != nil {
			return err
		}
		// fmt.Println(tagMode)
		goCode, err := GoObjectReference(tbName, structInstance, tagMode)
		if err != nil {
			return err
		}
		writeSqlFile(sqlFilename, goCode)
	}
	return nil
}

func writeSqlFile(filename string, sql string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(sql)
	return err
}

func batchDefault(tbName string, sqlFilename string, structInstance any) error {
	sql, err := mariadb.GenerateCreateTableSQLNeedTabName(tbName, structInstance)
	if err != nil {
		return err
	}
	sql += "\n"
	if err = writeSqlFile(sqlFilename, sql); err != nil {
		return err
	}
	sql, err = mariadb.GetSelectFieldsNeedTabName(tbName, structInstance)
	if err != nil {
		return err
	}
	sql += "\n"
	if err = writeSqlFile(sqlFilename, sql); err != nil {
		return err
	}
	sql, err = mariadb.GetUpdateFieldsNeedTabName(tbName, structInstance)
	if err != nil {
		return err
	}
	sql += "\n"
	err = writeSqlFile(sqlFilename, sql)
	return err
}

func batchAllFields(tbName string, sqlFilename string, structInstance any) error {
	sql, err := mariadb.GenerateCreateTableSQLAllFieldNeedTabName(tbName, structInstance)
	if err != nil {
		return err
	}
	sql += "\n"
	if err = writeSqlFile(sqlFilename, sql); err != nil {
		return err
	}
	sql, err = mariadb.GetSelectAllFieldNeedTabName(tbName, structInstance)
	if err != nil {
		return err
	}
	sql += "\n"
	if err = writeSqlFile(sqlFilename, sql); err != nil {
		return err
	}
	sql, err = mariadb.GetUpdateAllFieldNeedTabName(tbName, structInstance)
	if err != nil {
		return err
	}
	sql += "\n"
	err = writeSqlFile(sqlFilename, sql)
	return err
}

func batchCustomTag(tbName string, sqlFilename string, structInstance any, tagName string) error {
	sql, err := mariadb.GenerateCreateTableSQLCustomTagNeedTabName(tbName, structInstance, tagName)
	if err != nil {
		return err
	}
	sql += "\n"
	if err = writeSqlFile(sqlFilename, sql); err != nil {
		return err
	}
	sql, err = mariadb.GetSelectFieldsFromCustomTagNeedTabName(tbName, structInstance, tagName)
	if err != nil {
		return err
	}
	sql += "\n"
	if err = writeSqlFile(sqlFilename, sql); err != nil {
		return err
	}
	sql, err = mariadb.GetUpdateFieldsFromCustomTagNeedTabName(tbName, structInstance, tagName)
	if err != nil {
		return err
	}
	sql += "\n"
	err = writeSqlFile(sqlFilename, sql)
	return err
}
