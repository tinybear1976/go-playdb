package goplaydb

import "github.com/shopspring/decimal"

type Student struct {
	IsShow bool
	StuId  int             `json:"id" db:"id"`
	Name   string          `json:"name" db:"name"`
	Money  decimal.Decimal `json:"money" db:"money"`
}
