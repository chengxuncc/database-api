package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

type Response struct {
	Result  string      `json:"result"`
	Message interface{} `json:"message"`
}

func main() {
	var err error
	db, err = sql.Open("postgres", /*连接数据库*/
		fmt.Sprintf(
			"host=%s port=%s user=postgres dbname=%s password=%s sslmode=disable",
			HOST, PORT, DBNAME, PASSWORD,
		),
	)
	if err != nil { /*获取连接错误*/
		log.Panicln("Connect database Error:", err)
	}
	router := gin.Default()

	router.GET("/easy_select", func(c *gin.Context) {
		var goods []Goods /*切片*/
		var good Goods
		table_name := c.Query("table_name") /*接收查询的表*/
		Qualif := c.Query("qualif")         /*接收查询条件*/
		var query string
		if Qualif != "" {
			query = "select * from " + table_name + " where " + Qualif
		}
		if Qualif == "" {
			query = "select * from " + table_name
		}
		Rows, err := db.Query(query)
		if err == nil {
			if Rows != nil {
				for Rows.Next() {
					err = Rows.Scan(&good.Id, &good.Name, &good.Num, &good.Im_price, &good.Barcode_id, &good.Discount, &good.Sup_id)
					goods = append(goods, good) /*切片增加*/
				}
				Rows.Close()
				if err == nil {
					c.JSON(200, goods)
					return
				}
			}
		}
		c.JSON(200, Response{"error", err.Error()})
	})

	//router.GET("/Goods_query", func(c *gin.Context) {   /*按名字或条形码查询*/
	//	var goods []Goods  /*切片*/
	//	var good  Goods
	//	name := c.Query("name")      /*返回的string类型*/
	//	barcode_id :=c.Query("barcode_id")
	//	var query string
	//	if name != ""{
	//		query = "select * from goods where name = '"+name+"'"
	//	}
	//	if barcode_id != ""{
	//		query = "select * from goods where barcode_id = "+barcode_id
	//	}
	//	if name == ""&&barcode_id == ""{
	//		query = "select * from goods"
	//	}
	//	fmt.Println(query)
	//	Rows,err:= db.Query(query)
	//	if err ==nil{
	//		if Rows != nil{
	//			for Rows.Next(){
	//				err  =Rows.Scan(&good.Id,&good.Name,&good.Num,&good.Im_price,&good.Barcode_id,&good.Discount,&good.Sup_id)
	//				goods = append(goods, good)   /*切片增加*/
	//			}
	//			Rows.Close()
	//			if err == nil{
	//				c.JSON(200,goods)
	//				return
	//			}
	//		}
	//	}
	//	c.JSON(200,Response{"error",err.Error()})
	//})
	router.GET("/goods", GoodsGet)

	router.POST("/supplier", SupplierCreate)
	_ = router.Run()
}
