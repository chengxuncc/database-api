package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Sales struct {
	Id         string
	Gid        string
	Num        string
	Im_price   string
	Sale_price string
	Discount   string
	Time       string
	Paid_money string
	Employee   string
	Remark     string
}

func SalesGet(c *gin.Context) {
	var sales []Sales /*切片*/
	var asales Sales
	rows, err := db.Query("select * from goods")
	if err == nil {
		for rows.Next() {
			e := rows.Scan(getPointers(&asales)...)
			if e != nil {
				fmt.Println(e)
			}
			sales = append(sales, asales)
		}
		c.JSON(200, Response{"ok", sales})
		_ = rows.Close()
		return
	}
	c.JSON(200, Response{"error", err})
}
