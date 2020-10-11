package controllers

import (
	"DataCertProject/models"
	"astaxie/beego"
	"fmt"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post(){
	// 1. 解析数据
	var user models.User
	err :=r.ParseForm(&user)
	if err !=nil{
		r.Ctx.WriteString("解析错误，请重试！")
		return
	}
	r.Ctx.WriteString("123456")
	//2.保存信息到数据库

	id,err :=user.SaveUser()

	if err !=nil{
		r.Ctx.WriteString("失败")
		return
	}//注册成功

	r.TplName="login.html"
	fmt.Println(id)
	//3.返回前端结果


}
