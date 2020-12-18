package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func iniDB() (err error)  {

	dsn := "root:20206976@tcp(127.0.0.1:3306)/user"

	db,err = sql.Open("mysql", "root:20206976@tcp(127.0.0.1:3306)/user")

	if err != nil{

		fmt.Println("open %s failed,err:%v/n",dsn,err)
	}

	return

}


func main() {

	err := iniDB()
	if err != nil{
		fmt.Println("init DB failed,er:%v\n",err)
	}


	r := gin.Default()
	r.LoadHTMLFiles("./login.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})

	r.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		grade := c.PostForm("grade")
		email := c.PostForm("email")
		phone := c.PostForm("phone")
		teacher := c.PostForm("teacher")
		passGET6 := c.PostForm("passGET6")

		return name,grade,email,phone,teacher,passGET6

		})

	sqlStr := "insert into user(name, grade, email,phone, teacher,passGET6) values(name,grade,email,phone,teacher,passGET6)"



	r.Run(":8080")
    return
}
