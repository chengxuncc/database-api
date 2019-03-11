package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"reflect"
	"strings"
)

type Goods struct {
	Id         int
	Name       string
	Num        string
	Im_price   string
	Barcode_id string
	Discount   string
	Sup_id     string
}

type Supplier struct {
	Name    string
	Contact string
	Address string
}

type Response struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

func getFieldsValues(obj interface{}) (structname string, field string, value string) {  /*获取类型名等*/

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

func Insert(db *sql.DB, obj interface{}) error {
	tablename, fields, values := getFieldsValues(obj)
	sqls := fmt.Sprintf("insert into %s (%s) values (%s)", tablename, fields, values)
	_, err := db.Query(sqls)
	return err
}


func main() {
	db, err := sql.Open("postgres",     /*连接数据库*/
		fmt.Sprintf(
			"host=%s port=%s user=postgres dbname=%s password=%s sslmode=disable",
			"127.0.0.1","5432","postgres","312030",
		),
	)
	if err != nil {    /*获取连接错误*/
		log.Panicln("Connect database Error:", err)
	}
	router := gin.Default()

	router.GET("/easy_select", func(c *gin.Context) {
		var goods []Goods  /*切片*/
		var good  Goods
		table_name := c.Query("table_name")    /*接收查询的表*/
		Qualif := c.Query("qualif")     /*接收查询条件*/
		var query string
		if Qualif != ""{
			query = "select * from "+table_name+" where "+Qualif
		}
		if Qualif == ""{
			query = "select * from "+table_name
		}
		Rows,err:= db.Query(query)
		if err ==nil{
			if Rows != nil{
				for Rows.Next(){
					err  =Rows.Scan(&good.Id,&good.Name,&good.Num,&good.Im_price,&good.Barcode_id,&good.Discount,&good.Sup_id)
					goods = append(goods, good)   /*切片增加*/
				}
				Rows.Close()
				if err == nil{
					c.JSON(200,goods)
					return
				}
			}
		}
		c.JSON(200,Response{"error",err.Error()})
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

	router.POST("/supplier", func(c *gin.Context) {

		var o Supplier
		e := c.ShouldBindJSON(&o)
		if e == nil {
			e = Insert(db, o)
			if e == nil {
				c.JSON(200, gin.H{"result": "ok", "message": ""})
				return
			}
		}
		c.JSON(200, Response{"error", e.Error()})
	})
	_ = router.Run()
}
