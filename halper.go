package pgxscan

import (
	"errors"
	"reflect"
)

// create fields
func CreateFields(val any) (map[string]any, error) {

	v := reflect.ValueOf(val).Elem()
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return nil, errors.New("Неверный тип данных")
	}
	res := map[string]any{}
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanSet() && v.Field(i).CanAddr() {
			if vl, ok := t.Field(i).Tag.Lookup("db"); ok {
				if vl != "" {
					res[vl] = v.Field(i).Addr().Interface()
				}
			}
		}
	}
	return res, nil
}
