package pgxscan

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var cn *pgxpool.Pool = nil

// Инициализация подключения
func InitConnection(connectionstring string) error {
	var err error = nil
	cn, err = pgxpool.New(context.Background(), connectionstring)
	if err != nil {
		fmt.Println("Ошибка подключения к бд")
	}
	return err
}

// Сканирование запроса
func Scan[T any](sql string, args ...any) ([]T, error) {
	if cn == nil {
		return nil, errors.New("pool is not initialized")
	}
	rw, err := cn.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	res, err := ScanRows[T](rw)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ["map[bigint:2 byte:<nil> char:4 date:2026-02-05 00:00:00 +0000 UTC id:1 int:1 intarr:[1 2 3] json:map[ggg:rrrr] jsonb:map[ggg:rrrr] smallint:3 text:texttest time:2026-02-05 00:29:06.031303 +0000 UTC varchar:3]"]
// integer - int32
// biginteger - int64
// timestamp wotz - time.Time
// text - tring
// jsonb - map[string]any
// smallint - int16
// bytea - interface{}
// serial - int32
// character varying - string
// char - string
// date - time.Time
// json - map[string]interface{}
// integer[] - []interface{}

func ScanRows[T any](row pgx.Rows) ([]T, error) {
	val := *new(T)
	f, err := CreateFields(&val)
	if err != nil {
		return nil, err
	}
	res := *new([]T)
	for row.Next() {
		m, err := pgx.RowToMap(row)
		if err != nil {
			return nil, err
		}
		for key := range f {
			if vl, ok := m[key]; ok {
				switch vl.(type) {
				case int32: //int <- int32
					p := f[key].(*int)
					*p = int(vl.(int32))
				case time.Time: //tme <- time
					*f[key].(*time.Time) = vl.(time.Time)
				case map[string]any: // map[string]any <- json
					*f[key].(*map[string]any) = vl.(map[string]any)
				case string:
					*f[key].(*string) = vl.(string)
				default:
					return nil, fmt.Errorf("Неизвестный формат данных")
				}
			}
		}
		res = append(res, val)
	}
	return res, nil
}

func Query[T any](pgpool *pgxpool.Pool, sql string, args ...any) ([]T, error) {
	rw, err := pgpool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	res, err := ScanRows[T](rw)
	if err != nil {
		return nil, err
	}
	return res, nil
}
