package main

import "github.com/gin-gonic/gin"

type Supplier struct {
	Id      string
	Name    string
	Contact string
	Address string
	Time    string
}

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
