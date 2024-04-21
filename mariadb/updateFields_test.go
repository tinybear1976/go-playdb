package mariadb

import (
	"testing"
)

// type db1 struct {
// 	Id   int
// 	Name string `db:"username"`
// 	Age  int    `json:"ageage"`
// }
// type db2 struct {
// 	Id   int    `db:"userid"`
// 	Name string `db:"username"`
// 	Age  int
// }
// type a1 struct {
// 	Id   int
// 	Name string `axis:"A4"`
// 	Age  int    `axis:"B10"`
// }
// type a2 struct {
// 	Id   int    `axis_y:"4"`
// 	Name string `axis_y:"5"`
// 	Age  int
// }

func TestGetUpdateFields(t *testing.T) {
	testCases := []any{
		db1{},
		db2{},
		a1{},
		a2{},
	}
	for _, tc := range testCases {
		s, _ := GetUpdateFields(tc)
		t.Log(s)
	}
}

func TestGetUpdateAllField(t *testing.T) {
	testCases := []any{
		db1{},
		db2{},
		a1{},
		a2{},
	}
	for _, tc := range testCases {
		s, _ := GetUpdateAllField(tc)
		t.Log(s)
	}
}

func TestGetUpdateFieldsFromCustomTag(t *testing.T) {
	testCases := []any{
		db1{},
		db2{},
		a1{},
		a2{},
	}
	for _, tc := range testCases {
		s, _ := GetUpdateFieldsFromCustomTag(tc, "db")
		t.Log(s)
	}
}
