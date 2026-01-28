package pgxscan

import (
	"errors"
	"reflect"
)

// create fields
func CreateFields(val any) (map[string]any, error) {
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("Неверный тип данных")
	}
	res := map[string]any{}
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		if vl, ok := ft.Tag.Lookup("db"); ok {
			if vl != "" {
				res[vl] = v.Addr().Interface()
			}
		}
	}
	return res, nil
}
