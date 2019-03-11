package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func getFieldsValues(obj interface{}) (structname string, field string, value string) { /*获取类型名等*/

	t := reflect.TypeOf(obj)
	tv := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		tv = tv.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Panicln("Check type error not Struct")
		return
	}
	structname = t.Name()
	fieldNum := t.NumField()
	for i := 0; i < fieldNum-1; i++ {
		field += `"` + strings.ToLower(t.Field(i).Name) + `",`
		value += `'` + tv.Field(i).String() + `',`
	}
	field += t.Field(fieldNum - 1).Name
	value += `'` + tv.Field(fieldNum-1).String() + `'`
	return

}

//取插入对象的表名,字段名,值。值为空时则数据库默认
func getInsertParams(obj interface{}) (structname string, fields string, values string) {
	t := reflect.TypeOf(obj)
	tv := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		tv = tv.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Panicln("Check type error not Struct")
		return
	}
	structname = t.Name()
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		fields += `"` + strings.ToLower(t.Field(i).Name) + `"`
		if tv.Field(i).String() == "" {
			values += `default`
		} else {
			values += `'` + tv.Field(i).String() + `'`
		}
		if i != fieldNum-1 {
			fields += `,`
			values += `,`
		}
	}
	return
}

func Insert(db *sql.DB, obj interface{}) error {
	tablename, fields, values := getInsertParams(obj)
	sqls := fmt.Sprintf("insert into %s (%s) values (%s)", tablename, fields, values)
	_, err := db.Query(sqls)
	return err
}

func getPointers(obj interface{}) (res []interface{}) {
	tv := reflect.ValueOf(obj).Elem()
	fieldNum := tv.NumField()
	for i := 0; i < fieldNum; i++ {
		res = append(res, tv.Field(i).Addr().Interface())
	}
	return
}
