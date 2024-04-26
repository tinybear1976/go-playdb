package mariadb

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestGenerateCreateDatabaseSQL(t *testing.T) {
	testCases := []string{
		"dbName1",
		"dbName2",
	}

	for _, tc := range testCases {
		sql := GenerateCreateDatabaseSQL(tc)
		t.Logf("%s, Generated SQL: %v", tc, sql)
	}
}

func TestAst(t *testing.T) {

	txt := `package n
	  var n int
		type UserInfo struct {
		Id   int    ` + "`" + `json:"id" db:"id"` + "`" + "\nName string `" + `json:"name" db:"name"` + "`" + `}`

	// t.Log(txt)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", txt, 0)
	if err != nil {
		t.Error(err)
	}
	ast.Print(fset, f)

	type StructInfo struct {
		FieldName string
		FieldType string
		Tag       string
	}
	objects := make(map[string][]StructInfo)
	var tmp []StructInfo = nil
	var tmpName string = ""
	ast.Inspect(f, func(x ast.Node) bool {
		typ, typOk := x.(*ast.TypeSpec)
		if typOk {
			// fmt.Println("type name:", typ.Name.Name)
			objects[typ.Name.Name] = nil
			if tmp != nil {
				objects[tmpName] = tmp
			}
			tmpName = typ.Name.Name
			return true
		}
		s, ok := x.(*ast.StructType)
		if !ok {
			return true
		}
		if tmp == nil {
			tmp = []StructInfo{}
		}
		for _, field := range s.Fields.List {
			//fmt.Printf("%#v", field)
			// fmt.Println("field.Type:", field.Type)
			_st := StructInfo{
				FieldName: field.Names[0].Name,
				FieldType: fmt.Sprintf("%v", field.Type),
				Tag:       field.Tag.Value,
			}
			tmp = append(tmp, _st)
			// fmt.Printf("Field: %s\n", field.Names[0].Name)
			// fmt.Printf("Tag:   %s\n", field.Tag.Value)
		}
		return false
	})
	if tmp != nil {
		objects[tmpName] = tmp
	}
	fmt.Printf("%#v\n", objects)

	ttt := reflect.StructOf([]reflect.StructField{
		{
			Name: "A",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"a"`,
		},
		{
			Name: "B",
			Type: reflect.TypeOf(""),
			Tag:  `json:"B"`,
		},
	})
	fmt.Println("struct name:", ttt.Name())

}
