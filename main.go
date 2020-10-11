package main

import (
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 1.链接数据库
	db_mysql.ConnectDB()

	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()
}

