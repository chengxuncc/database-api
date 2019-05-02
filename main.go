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
	db, err = sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%s user=postgres dbname=%s password=%s sslmode=disable",
			"127.0.0.1", "5432", "postgres", "312030",
		),
	)
	if err != nil {
		log.Panicln("Connect database Error:", err)
	}
	router := gin.Default()

	router.GET("/goods", GoodsGet)
	router.GET("/goods/:search", GoodsSearch)
	router.GET("/goods/remove", GoodsDel)

	router.GET("/employee", EmployeeGet)

	router.POST("/supplier", SupplierCreate)
	router.GET("/supplier", Sups_Get)
	_ = router.Run()
}
