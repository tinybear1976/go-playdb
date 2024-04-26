package goplaydb

type Student struct {
	IsShow bool
	StuId  int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
}
