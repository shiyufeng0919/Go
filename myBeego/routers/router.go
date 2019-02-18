package routers

import (
	"myBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego vs gin示例
	//浏览器请求:http://127.0.0.1:8083/detail/123
    //beego.Router("/detail/?:id", &controllers.MainController{})

    //beego搭建web框架示例
    beego.Include(&controllers.UserController{})
}
