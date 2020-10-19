package controllers

import (
	"DataCertProject/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"log"
	"os"
	"path"
	"time"
)

type UploadController struct{
	beego.Controller
}

func (this *UploadController) Get() {
	this.TplName="upload.html"     //显示开始界面
}

func (this *UploadController)Post() {

	phone := this.Ctx.Request.PostFormValue("phone")

	file, information, err := this.GetFile("file")  //返回文件，文件信息头，错误信息
	if err != nil {
		this.Ctx.WriteString("File retrieval failure")
		return
	}

	fmt.Println("文件名称",information.Filename)
	a :=information.Size /1024
	fmt.Println("文件大小", a,"KB") //字节大小
	defer file.Close()    //关闭上传的文件，否则出现临时文件不清除的情况

	filename := information.Filename           //将文件信息头的信息赋值给filename变量
	err = this.SaveToFile("file", path.Join("static/upload",filename))  //保存文件的路径。保存在static/upload中 （文件名）

	file,err  =os.Open("./static/upload/"+filename)

	defer file.Close()
	if err !=nil{
		fmt.Println(err.Error())
		return
	}
	hash :=md5.New()
	if _,err :=io.Copy(hash,file) ; err !=nil{
		log.Fatal(err.Error())
	}
	sum :=hash.Sum(nil)
	sums:=hex.EncodeToString(sum)
	fmt.Printf("%x\n",sum)


	if err != nil {
		this.Ctx.WriteString("File upload failed！")
	}
	//} else {
	//	this.TplName = "sgin_in.html"
	//	return    //上传成功后显示信息
	//}
	record :=models.Upload{}
	record.FileName =information.Filename
	record.FileSize = information.Size
	record.CertTime =time.Now().Unix()
	record.FileCert = sums
	record.Phone = phone

	_,err = record.Saveupload()
	if err !=nil{
		this.Ctx.WriteString("数据认证错误")
	}

	recores,err :=models.QueryPhone(phone)
	if err !=nil{
		this.Ctx.WriteString("获取认证数据失败")
	}
	fmt.Println(recores)
	this.Data["Records"] = recores
	this.Data["Phone"] = phone
	
	this.TplName = "record.html"             //停留在当前界面
}
