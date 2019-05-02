package main

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupplier_Create(t *testing.T) {
	db, err := sql.Open("postgres", /*连接数据库*/
		fmt.Sprintf(
			"host=%s port=%s user=postgres dbname=%s password=%s sslmode=disable",
			HOST, PORT, DBNAME, PASSWORD,
		),
	)
	assert.Equal(t, nil, err)
	err = Insert(db, Supplier{Id:"", Name:"Test_供应商名", Contact:"联系方式", Address:"地址", Add_time:""})
	assert.Equal(t, nil, err)
}
