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
	err = Insert(db, Supplier{"", "Test_供应商名", "联系方式", "地址", ""})
	assert.Equal(t, nil, err)
}
