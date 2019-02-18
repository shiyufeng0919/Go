package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myBeego/models"
)

type MainController struct {
	beego.Controller
}

//模拟通过id从数据库查询相应的信息
func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	//假装提取了id
	id:= c.Ctx.Input.Param(":id")
	fmt.Printf("id is:%s\n",id)

	//设定模版
	c.TplName = "index.html"

	//向模版绑定数据,用于前端提取数据
	c.Data["Photos"]=models.GetPhotos()
	c.Data["Title"]=models.GetTitle()
	c.Data["Price"]=models.GetPriceString()
}
