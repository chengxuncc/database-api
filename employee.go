package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*获取职员信息*/
type Employee struct {
	Id         string
	Name       string
	Age        string
	Sex        string
	Position   string
	Salary     string
	Contact    string
	Site       string
	Entrytime  string
	Resigntime string
	Is_delete  bool
}

func EmployeeGet(c *gin.Context) {
	var Employees []Employee /*切片*/
	var aEmployee Employee
	rows, err := db.Query("select * from employee")
	if err == nil {
		for rows.Next() {
			e := rows.Scan(getPointers(&aEmployee)...)
			if e != nil {
				fmt.Println(e)
			}
			Employees = append(Employees, aEmployee)
		}
		c.JSON(200, Response{"ok", Employees})
		return
	}
	err = rows.Close()
	c.JSON(200, Response{"error", err})
}

func EmployeeRm(c *gin.Context) {
	remove(c, "employee")
}
