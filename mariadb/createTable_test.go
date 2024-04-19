package mariadb

import (
	"reflect"
	"testing"
)

func TestGenerateCreateTableSQL(t *testing.T) {
	type S1 struct {
		Name string `axis:"A4"`
		Age  int    `axis:"B4"`
	}
	type S2 struct {
		Name string `axis_y:"5"`
		Age  int    `axis_y:"6"`
	}
	testCases := []any{
		S1{},
		S2{},
		"string",
	}

	for _, tc := range testCases {
		sql, err := GenerateCreateTableSQL(tc)
		typName := reflect.TypeOf(tc).Name()
		if err != nil {
			t.Logf("%s error: %v", typName, err)
		} else {

			t.Logf("%s, Generated SQL: %v", typName, sql)
		}
	}
}
