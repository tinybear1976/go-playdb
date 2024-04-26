package goplaydb

import "testing"

func TestReadGoSourceFile(t *testing.T) {
	content, err := readGoSourceFile("test_go.go")
	if err != nil {
		t.Error(err)
	}
	t.Log(content)
}
func TestGetStructs(t *testing.T) {
	objs, err := GetStructs("test_go.go")
	if err != nil {
		t.Error(err)
	}
	t.Log(objs)
}
func TestGenerateSQLFromFile(t *testing.T) {
	err := GenerateSQLFromFile("test_go.go", "")
	if err != nil {
		t.Log(err)
	}
	t.Log("GenerateSQLFromFile success")
}

func TestGenerateSQLFromFileAll(t *testing.T) {
	err := GenerateSQLFromFile("test_go.go", "*")
	if err != nil {
		t.Log(err)
	}
	t.Log("GenerateSQLFromFile success")
}

func TestGenerateSQLFromFileCustom(t *testing.T) {
	err := GenerateSQLFromFile("test_go.go", "db")
	if err != nil {
		t.Log(err)
	}
	t.Log("GenerateSQLFromFile success")
}
