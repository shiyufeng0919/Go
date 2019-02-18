package main

import (
	"github.com/astaxie/beego/orm"
	_ "myBeego/routers" //_代表只执行routers目录下的init方法
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//beego搭建web框架示例
	//数据库注册 用户名:密码@数据库地址+名称?字符集
	orm.RegisterDataBase("default","mysql","root:shiyufeng@tcp(127.0.0.1:3306)/learn")

	beego.Run()
}

