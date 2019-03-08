package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"reflect"
)

type Goods struct {
	name       string
	num        int
	im_price   float32
	barcode_id int
	discount   float32
	sup_id     int
}

type supplier struct {
	name string
	contact string
	address string
}

func getFieldValue(structName interface{}) (structname string,field string,value string)  {

	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Panicln("Check type error not Struct")
		return
	}
	structname=t.Name()
	tv:=reflect.ValueOf(structName)
	fieldNum := t.NumField()
	for i:= 0;i<fieldNum-1;i++ {
		field+=t.Field(i).Name+","
		value+=`'`+tv.Field(i).String()+`',`
	}
	field+=t.Field(fieldNum-1).Name
	value+=`'`+tv.Field(fieldNum-1).String()+`'`
	return

}

func Insert(db *sql.DB,obj interface{}) error  {
	tablename,fields,values:=getFieldValue(obj)
	_,err:=db.Query("insert into "+tablename+" ("+fields+") values ("+values+")")
	return err
}

func main() {
	db, err := sql.Open("postgres", "host=172.19.153.61 port=5432 user=postgres dbname=postgres password=a123456 sslmode=disable")
	if err!=nil{
		log.Panicln("Connect database Error:",err)
	}
	s:=supplier{
		name:"中粮2",
		contact:"1380013880000",
		address:"安徽省宣城市",
	}
	err = Insert(db,s)
	fmt.Println("supplier err:", err)
	os.Exit(0)
	g:=Goods{
		name:       "可乐",
		num:        10,
		im_price:   3.5,
		barcode_id: 11111,
		discount:   0,
		sup_id:     3,
	}
	_,err = db.Query("insert into goods (name,num,im_price,barcode_id,discount,sup_id) values ($1,$2,$3,$4,$5,$6)",
		g.name,
		g.num,
		g.im_price,
		g.barcode_id,
		g.discount,
		g.sup_id,
		)
	fmt.Println("Goods err:", err)
	router := gin.Default()
	router.POST("/form_post", func(c *gin.Context) {

		message := c.PostForm("message")

		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{"status": gin.H{"status_code":

		http.StatusOK, "status": "ok",}, "message": message, "nick": nick,
		})
	})
	router.Run();
}
