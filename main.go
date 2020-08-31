package main

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// DemoRow for `test_table`
type DemoRow struct {
	FieldKey string  `sql:"field_key"`
	FieldOne string  `sql:"field_one"`
	FieldTwo bool    `sql:"field_two"`
	FieldThr int64   `sql:"field_thr"`
	FieldFou float64 `sql:"field_fou"`
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root123@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	if db == nil {
		fmt.Println("db is nil.")
		return
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)

	db.SetConnMaxLifetime(time.Hour)
	db.Ping()

}
func main() {

	table := "test_table"

	// test Query

	rowArr1, _ := QueryByFieldOne(table, nil, "0", "field_two")
	rowArrAll, _ := QueryAll(table, nil)
	fmt.Println(rowArr1)
	fmt.Println(rowArrAll)

	// test Update

	tx, _ := db.Begin()

	row0, _ := QueryByKey(table, tx, "key001")
	row0.FieldOne = "one123"
	row0.FieldTwo = true
	row0.FieldThr = 123455
	row0.FieldFou = 123.45
	_ = Update(table, tx, row0)

	tx.Commit()

	// test Insert
	newRow0 := DemoRow{
		FieldKey: "key002",
		FieldOne: "one456",
		FieldTwo: false,
		FieldThr: 5678,
		FieldFou: 0.01,
	}
	newRow1 := DemoRow{
		FieldKey: "key003",
		FieldOne: "one789",
		FieldTwo: true,
		FieldThr: 5678,
		FieldFou: 0.02,
	}

	_ = Insert(table, nil, newRow0, newRow1)

	// test Remove
	_ = Remove(table, nil, "key011")

	fmt.Println("====>End")
}

// Query by primary key (field[0])
func QueryByKey(table string, tx *sql.Tx, fieldKey string) (
	*DemoRow, error) {
	ctx := context.Background()
	var row DemoRow
	row.FieldKey = fieldKey
	fm, err := NewFieldsMap(table, &row)
	if err != nil {
		return nil, err
	}

	objptr, err := fm.SQLSelectByPriKey(ctx, tx, db)
	if err != nil {
		return nil, err
	}
	fmt.Println(reflect.TypeOf(objptr))

	return objptr.(*DemoRow), nil
}

// Query by `field_one`
func QueryByFieldOne(table string, tx *sql.Tx, fieldOne string, field string) (
	[]DemoRow, error) {
	ctx := context.Background()
	var row DemoRow
	row.FieldOne = fieldOne
	fm, err := NewFieldsMap(table, &row)
	if err != nil {
		return nil, err
	}

	objptrs, err := fm.SQLSelectRowsByFieldNameInDB(ctx, tx, db, field)
	if err != nil {
		return nil, err
	}

	var objs []DemoRow
	for i, olen := 0, len(objptrs); i < olen; i++ {
		objs = append(objs, *objptrs[i].(*DemoRow))
	}

	return objs, nil
}

// Query all
func QueryAll(table string, tx *sql.Tx) ([]DemoRow, error) {
	ctx := context.Background()
	var row DemoRow
	fm, err := NewFieldsMap(table, &row)
	if err != nil {
		return nil, err
	}

	objptrs, err := fm.SQLSelectAllRows(ctx, tx, db)
	if err != nil {
		return nil, err
	}

	var objs []DemoRow
	for i, olen := 0, len(objptrs); i < olen; i++ {
		objs = append(objs, *objptrs[i].(*DemoRow))
	}

	return objs, nil
}

// Insert
func Insert(table string, tx *sql.Tx, rows ...DemoRow) error {
	ctx := context.Background()
	for i, tlen := 0, len(rows); i < tlen; i++ {

		fm, err := NewFieldsMap(table, &rows[i])
		if err != nil {
			return err
		}

		err = fm.SQLInsert(ctx, tx, db)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update by primary key (field[0])
func Update(table string, tx *sql.Tx, row *DemoRow) error {
	ctx := context.Background()
	fm, err := NewFieldsMap(table, row)
	if err != nil {
		return err
	}

	err = fm.SQLUpdateByPriKey(ctx, tx, db)
	if err != nil {
		return err
	}

	return nil
}

// Remove by primary key (field[0])
func Remove(table string, tx *sql.Tx, fieldKey string) error {
	ctx := context.Background()
	var row DemoRow
	row.FieldKey = fieldKey
	fm, err := NewFieldsMap(table, &row)
	if err != nil {
		return err
	}

	err = fm.SQLDeleteByPriKey(ctx, tx, db)
	if err != nil {
		return err
	}

	return nil
}
