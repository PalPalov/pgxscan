package pgxscan_test

import (
	"pgxscan"
	"testing"
)

type FieldTest struct {
	id   int `db:"id"`
	name string
}

func TestCreateFields(t *testing.T) {
	ft := FieldTest{}
	pgxscan.CreateFields(ft)
}
