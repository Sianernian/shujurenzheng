package models

import (
	"DataCertProject/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Pwd string `form:"password"`
}


func (u User) SaveUser()(int64 ,error){

	// 将用户密码进行hash ，使用md5 计算hash值
	hashMd5:= md5.New() // 实例化一个 hash
	hashMd5.Write([]byte(u.Pwd))
	bytes :=hashMd5.Sum(nil)
	u.Pwd = hex.EncodeToString(bytes)

	// 执行数据库

	row ,err:=db_mysql.Db.Exec("insert into user(Phone,pwd)  values(?,?)",u.Phone,u.Pwd)

	if err !=nil{
		return -1,err
	}
	id ,err :=row.RowsAffected()
	if err !=nil{
		return -1 ,err
	}
	return id , nil
}