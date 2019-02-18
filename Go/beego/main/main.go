package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

//1.定义RESTful Controller
type RESTfulController struct {
	beego.Controller
}
//定义GET方法
func (this *RESTfulController) Get(){
	this.Ctx.WriteString("hello beego in GET Method!")
}
//定义POST方法
func (this *RESTfulController) Post(){
	this.Ctx.WriteString("hello beego in POST Method!")
}


//2.正则路由
type RegExpController struct {
	beego.Controller
}
func (this *RegExpController) Get(){
	//1.用于提取id
	this.Ctx.WriteString(fmt.Sprintln("In RegExp Mode!"))
	id:=this.Ctx.Input.Param(":id")
	this.Ctx.WriteString(fmt.Sprintf("id is:%s \n",id))

	//2.用于提取*内容 postman->http://127.0.0.1:8081/RegExp5/learn/beego,结果:splat is: learn/beego
	splat:=this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString((fmt.Sprintf("splat is: %s\n",splat)))

	//3.提取beego.go中的beego和go
	path:=this.Ctx.Input.Param(":path")
	this.Ctx.WriteString((fmt.Sprintf("path is: %s\n",path))) //beego

	ext:=this.Ctx.Input.Param(":ext")
	this.Ctx.WriteString((fmt.Sprintf("ext is: %s\n",ext))) //go
}

/*
##RESTful方式
方式一：地址栏请求GET方法
  启动main方法，地址栏Get请求url: http://127.0.0.1:8081/RESTful
  访问RESTfulController的Get方法，页面打印:hello beego in GET Method!
方式二：postman请求post方法
  访问RESTfulController的Post方法,打印hello beego in POST Method!
##正则方式
*/
func main(){
	//RESTful Controller路由
	beego.Router("/RESTful",&RESTfulController{})

	//正则路由，从path中提取参数
	//1.提取id
	beego.Router("/RegExp1/?:id",&RegExpController{}) //id任意,postman请求:http://127.0.0.1:8081/RegExp1/beego,id=beego
	beego.Router("/RegExp2/:id([0-9]+)",&RegExpController{}) //id只能为0-9的数字,可以有多数字
	beego.Router("/RegExp3/:id([\\w]+)",&RegExpController{}) //id不能为空，\\w所有非空格字符
	beego.Router("/RegExp4/bee:id([0-9]+)go",&RegExpController{}) //以bee开头，go结尾，中间由若干数字构成，并可被提取(postman:http://127.0.0.1:8081/RegExp4/bee123go,提取结果123)
	//2.提取所有
	beego.Router("/RegExp5/*",&RegExpController{}) //用于提取RegExp5后的所有内容：postman-> http://127.0.0.1:8081/RegExp5/learn/beego
	//3.提取beego.go中的beego和go
	beego.Router("/RegExp6/*.*",&RegExpController{})//postman->http://127.0.0.1:8081/RegExp6/beego.go

	//服务启动
	beego.Run("127.0.0.1:8081")
}