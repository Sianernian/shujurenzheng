package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 注册 页面
    beego.Router("/", &controllers.MainController{})
    //登录界面
    beego.Router("/register",&controllers.RegisterController{})
    // 主页面
    beego.Router("/sgin_in",&controllers.SginController{})
    //文件上传接口
	beego.Router("/", &controllers.UploadController{},"*:Get")
	beego.Router("/home", &controllers.UploadController{},"*:Post")

    // 数据在列表页面
    beego.Router("/record",&controllers.UploadController{})
    //认证按钮 跳转 新增页面
    //beego.Router("/upload_file",)

}
