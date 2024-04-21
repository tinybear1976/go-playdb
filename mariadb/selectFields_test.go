package mariadb

import (
	"strings"
	"testing"
)

func TestStringJoin(t *testing.T) {
	testCases := [][]string{
		// 测试用例1
		{"id"},
		// 测试用例2
		{"id", "name", "age", "gender"},
		// 测试用例3
		{},
	}

	for _, tc := range testCases {
		t.Log(strings.Join(tc, ", "))
	}
}

type db1 struct {
	Id   int
	Name string `db:"username"`
	Age  int    `json:"ageage"`
}
type db2 struct {
	Id   int    `db:"userid"`
	Name string `db:"username"`
	Age  int
}
type a1 struct {
	Id   int
	Name string `axis:"A4"`
	Age  int    `axis:"B10"`
}
type a2 struct {
	Id   int    `axis_y:"4"`
	Name string `axis_y:"5"`
	Age  int
}

func TestGetSelectFields(t *testing.T) {
	testCases := []any{
		db1{},
		db2{},
		a1{},
		a2{},
	}
	for _, tc := range testCases {
		s, _ := GetSelectFields(tc)
		t.Log(s)
	}
}

func TestGetSelectAllField(t *testing.T) {
	testCases := []any{
		db1{},
		db2{},
		a1{},
		a2{},
	}
	for _, tc := range testCases {
		s, _ := GetSelectAllField(tc)
		t.Log(s)
	}
}

func TestGetSelectFieldsFromCustomTag(t *testing.T) {
	testCases := []any{
		db1{},
		db2{},
		a1{},
		a2{},
	}
	for _, tc := range testCases {
		s, _ := GetSelectFieldsFromCustomTag(tc, "db")
		t.Log(s)
	}
}
