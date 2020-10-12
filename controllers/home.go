package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"log"
	"os"
	"path"
)

type UploadController struct{
	beego.Controller
}

func (this *UploadController) Get() {
	this.TplName="upload.html"     //显示开始界面
}

func (this *UploadController)Post() {
	file, information, err := this.GetFile("file")  //返回文件，文件信息头，错误信息
	if err != nil {
		this.Ctx.WriteString("File retrieval failure")
		return
	}
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
	fmt.Printf("%x\n",sum)


	if err != nil {
		this.Ctx.WriteString("File upload failed！")
	} else {
		this.Ctx.WriteString("File upload succeed!")  //上传成功后显示信息
	}

	this.TplName = "sgin_in.html"             //停留在当前界面
}
