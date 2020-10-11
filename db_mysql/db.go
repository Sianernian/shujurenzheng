package db_mysql

import (
	"astaxie/beego"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	)
var Db *sql.DB

func ConnectDB(){

	config :=beego.AppConfig

	dbDriver :=config.String("db_driverName")
	dbUser :=config.String("db_user")
	dbPwd := config.String("db_pwd")
	dbIp := config.String("db_ip")
	dbPort :=config.String("da_pore")
	dbName :=config.String("da_name")

	conmint :=dbUser+":"+dbPwd+"@tcp("+dbIp+":"+dbPort+")/"+dbName+"?charset=utf8"
	fmt.Println(conmint)
	DB,err:=sql.Open(dbDriver,conmint)
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("成功")
	Db = DB
}

