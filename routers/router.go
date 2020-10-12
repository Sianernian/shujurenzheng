package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/sgin_in",&controllers.SginController{})
	beego.Router("/", &controllers.UploadController{},"*:Get")
	beego.Router("/home", &controllers.UploadController{},"*:Post")
}
