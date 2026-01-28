package pgxscan_test

import (
	"pgxscan"
	"slices"
	"testing"
)

var fields = []string{"id", "name"}

type FieldTest struct {
	id    int `db:"id"`
	name  string
	name2 string `db:"name"`
	name3 string `db:""`
}

func TestCreateFields(t *testing.T) {
	ft := FieldTest{}
	flist, err := pgxscan.CreateFields(ft)
	if err != nil {
		t.Error("Неверный возврат")
	}
	if !slices.Equal(flist, fields) {
		t.Errorf("Неверное значение возврата функции %v", flist)
	}

}
