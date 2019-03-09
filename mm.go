package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Goods struct {
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

func getFieldsValues(obj interface{}) (structname string, field string, value string) {

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
	fmt.Println(sqls)
	_, err := db.Query(sqls)
	return err
}

func main() {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%s user=postgres dbname=%s password=%s sslmode=disable",
			HOST, PORT, DBNAME, PASSWORD,
		),
	)
	if err != nil {
		log.Panicln("Connect database Error:", err)
	}
	router := gin.Default()
	router.POST("/form_post", func(c *gin.Context) {

		message := c.PostForm("message")

		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{"status": gin.H{"status_code": http.StatusOK, "status": "ok"}, "message": message, "nick": nick,
		})
	})

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
