package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type SginController struct{
	beego.Controller
}


// 刷新
func (s *SginController)Get(){
	s.TplName = "sgin_in.html"
}

func (s *SginController) Post(){
	var user models.User

	err :=s.ParseForm(&user)

	if err !=nil{
		s.Ctx.WriteString("解析错误，请重试！")
		return
	}

	a,err := user.Query()
	if err !=nil{
		fmt.Println(err.Error())

		return
	}
	s.Data["Phone"] =a.Phone

	s.TplName="sgin_in.html" //{{.Phone}}
}