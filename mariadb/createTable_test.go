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

func TestGenerateCreateTableSQAllField(t *testing.T) {
	type S1 struct {
		Id   int
		Name string
		Age  int `json:"ageage"`
	}
	type S2 struct {
		Id   string
		Name string
		Age  int
	}
	testCases := []any{
		S1{},
		S2{},
	}

	for _, tc := range testCases {
		sql, err := GenerateCreateTableSQLAllField(tc)
		typName := reflect.TypeOf(tc).Name()
		if err != nil {
			t.Logf("%s error: %v", typName, err)
		} else {

			t.Logf("%s, Generated SQL: %v", typName, sql)
		}
	}
}

func TestGenerateCreateTableSQLCustomTag(t *testing.T) {
	type S1 struct {
		Id   int
		Name string `db:"username"`
		Age  int    `json:"ageage"`
	}
	type S2 struct {
		Id   int    `db:"userid"`
		Name string `db:"username"`
		Age  int
	}
	testCases := []any{
		S1{},
		S2{},
	}

	for _, tc := range testCases {
		sql, err := GenerateCreateTableSQLCustomTag(tc, "db")
		typName := reflect.TypeOf(tc).Name()
		if err != nil {
			t.Logf("%s error: %v", typName, err)
		} else {

			t.Logf("%s, Generated SQL: %v", typName, sql)
		}
	}
}
