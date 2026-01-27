package pgxscan

import (
	"errors"
	"fmt"
	"reflect"
)

// create fields
func CreateFields(val any) ([]string, error) {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Struct {
		return nil, errors.New("Неверный тип данных")
	}

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		fv.Type().Field(i).tag
	}

	fmt.Println(v.Kind())
	return nil, nil
}
