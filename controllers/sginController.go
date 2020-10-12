package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type SginController struct{
	beego.Controller
}


func (s *SginController) Post(){
	var user models.User
	err :=s.ParseForm(&user)
	if err !=nil{
		s.Ctx.WriteString("解析错误，请重试！")
		return
	}

	err = user.Query()
	if err !=nil{
		s.TplName = "sgin_in.html"
		return
	}
	s.TplName="sgin_in.html"
}