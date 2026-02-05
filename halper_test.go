package pgxscan_test

import (
	"fmt"
	"pgxscan"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var fields = []string{"id", "name"}

type FieldTest struct {
	Id    int `db:"id"`
	Name  string
	Name2 string `db:"name"`
	Name3 string `db:""`
}

type TestStruct struct {
	Vl1  int    `db:"id"`
	Vl2  string `db:"int"`
	Vl3  string `db:"bigint"`
	Vl4  string `db:"varchar"`
	Vl5  string `db:"char"`
	Vl6  string `db:"time"`
	Vl7  string `db:"date"`
	Vl8  string `db:"text"`
	Vl9  string `db:"json"`
	Vl10 string `db:"jsonb"`
}

const connstring = "postgres://masteruser:GfHtrj&6g@10.0.3.78:5432/mng_data"

func TestCreateFields(t *testing.T) {
	pool, err := pgxpool.New(t.Context(), connstring)
	if err != nil {
		t.Error("Ошибка подключения к бд")
	}
	rw, err := pool.Query(t.Context(), "select * from rqts")
	if err != nil {
		t.Error(err.Error())
	}
	//rr := &rw
	//ft := FieldTest{}
	ft, err := pgxscan.Scan[FieldTest](rw)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(ft)
	// ft := FieldTest{}
	// flist, err := pgxscan.CreateFields(ft)
	// if err != nil {
	// 	t.Error("Неверный возврат")
	// }

	// if !slices.Equal(flist, fields) {
	// 	t.Errorf("Неверное значение возврата функции %v", flist)
	// }

}

type MyStruct struct {
	Id int
}

func TestF(tst *testing.T) {
	ms := &MyStruct{}
	//t := reflect.TypeOf(ms)
	v := reflect.ValueOf(ms).Elem()
	for i := 0; i < v.NumField(); i++ {

		field := v.Field(i)
		if field.CanAddr() {
			vl := field.Addr().Interface()
			vlp := vl.(*int)
			*vlp = 10
			fmt.Print(vl)
		}
	}
	fmt.Print(ms)
}
