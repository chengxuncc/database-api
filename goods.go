package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goods struct {
	Id         string
	Name       string
	Num        string
	Im_price   string
	Barcode_id string
	Discount   string
	Sup_id     string
	Time       string
}

func GoodsGet(c *gin.Context) {
	var goods []Goods /*切片*/
	var aGoods Goods
	rows, err := db.Query("select * from goods")
	if err == nil {
		for rows.Next() {
			e := rows.Scan(getPointers(&aGoods)...)
			if e != nil {
				fmt.Println(e)
			}
			goods = append(goods, aGoods)
		}
		c.JSON(200, Response{"ok", goods})
		_ = rows.Close()
		return
	}
	c.JSON(200, Response{"error", err})
}

func GoodsSearch(c *gin.Context) {

	searchString:=c.Param("search")
	var goods []Goods /*切片*/
	var aGoods Goods
	rows, err := db.Query(`select * from goods where "name" like '%?%' or "barcode_id" like '%?%'`,
		searchString,searchString)
	if err == nil {
		for rows.Next() {
			e := rows.Scan(getPointers(&aGoods)...)
			if e != nil {
				fmt.Println(e)
			}
			goods = append(goods, aGoods)
		}
		c.JSON(200, Response{"ok", goods})
		_ = rows.Close()
		return
	}
	c.JSON(200, Response{"error", err})
}
