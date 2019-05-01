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
	Add_time   string
	Upd_time   string
	Del_time   string
	Is_delete  bool
	Price      string
}

/*获取商品列表*/
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

/*插入商品信息*/
func GoodsCreate(c *gin.Context) {

	var o Goods
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

/*商品删除*/
func GoodsDel(c *gin.Context) {
	remove(c, "goods")
}
