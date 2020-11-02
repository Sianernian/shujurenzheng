package models

import (
	"DataCertProject/db_mysql"
	"DataCertProject/util"
	"fmt"
)

type Upload struct {
	Id int
	FileName string
	FileSize int64
	FileCert string // 认证号
	CertTime int64
	FormatCertTime string //格式化时间格式，该字段仅在前端展示时使用
	Phone string
}

/*
 * 存入上传数据到数据库上
 */

func (u Upload) Saveupload()(int64 ,error){
	row ,err:=db_mysql.Db.Exec("insert into upload(file_name ,file_size , file_cert , cert_time,phone) " +
		"values(?,?,?,?,?)",
		u.FileName,
		u.FileSize,
		u.FileCert,
		u.CertTime,
		u.Phone)
	if err !=nil{
		fmt.Println(err.Error())
		return -1,err
	}
	id ,err :=row.RowsAffected()
	if err !=nil{
		return -1,err
	}
	return id,nil
}
func QueryPhone(phone string)([]Upload ,error){
	rs ,err :=db_mysql.Db.Query("select id ,file_name ,file_size,file_cert, cert_time,phone from upload where phone =?",phone)
	if err !=nil{
		return nil,err
	}
	records :=make([]Upload,0)
	for rs.Next(){
		var record Upload
		err := rs.Scan(&record.Id,&record.FileName,&record.FileSize , &record.FileCert,&record.CertTime,&record.Phone)
		if err !=nil{
			return  nil,err
		}
		//时间转换
		record.FormatCertTime = util.TimeFormat(record.CertTime,0,util.TIME_FORMAT_THREE)
		records = append(records, record)
	}
	return records, nil
}