package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Import struct {
	Id   string
	Num  string
	Gid  string
	Time string
}

func ImportGet(c *gin.Context) {
	var imports []Import /*切片*/
	var aimport Import
	rows, err := db.Query("select * from 'import'")
	if err == nil {
		for rows.Next() {
			e := rows.Scan(getPointers(&aimport)...)
			if e != nil {
				fmt.Println(e)
			}
			imports = append(imports, aimport)
		}
		c.JSON(200, Response{"ok", imports})
		_ = rows.Close()
		return
	}
	c.JSON(200, Response{"error", err})
}
