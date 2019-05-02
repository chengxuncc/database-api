package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Supplier struct {
	Id        string
	Name      string
	Contact   string
	Address   string
	Add_time  string
	Upd_time  string
	Del_time  string
	Is_delete bool
}

/*获取供货商列表*/
func Sups_Get(c *gin.Context) {
	var suppliers []Supplier
	var asupplier Supplier
	rows, err := db.Query("select * from supplier")
	if err == nil {
		for rows.Next() {
			e := rows.Scan(getPointers(&asupplier)...)
			if e != nil {
				fmt.Println(e)
			}
			suppliers = append(suppliers, asupplier)
		}
		c.JSON(200, Response{"ok", suppliers})
		return
	}
	err = rows.Close()
	c.JSON(200, Response{"error", err})
}

/*新建新的供货商*/
func SupplierCreate(c *gin.Context) {

	var o Supplier
	e := c.ShouldBindJSON(&o)
	if e == nil {
		e = Insert(db, o)
		if e == nil {
			c.JSON(200, Response{"ok", ""})
			return
		}
	}
	c.JSON(200, Response{"error", e.Error()})
}
