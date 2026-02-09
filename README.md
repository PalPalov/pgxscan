# pgxscan

## Description

This module supports the use of the pgx library. A scan method has been implemented that returns a slice of certain structures.

## Usage

To load data into the structure fields, you must define a db tag indicating the name of the field in the database.

```
type FieldTest struct {
	Id    int `db:"id"`
	Name string `db:"name"`
}

func ReadFunction(){
    pgxscan.InitConnection(CONNECTIONSTRING)
	ft, err := pgxscan.Scan[FieldTest](A request containing the id and name fields)
	if err != nil {
		panic((err.Error()))
	}

	fmt.Println(ft)
}
```