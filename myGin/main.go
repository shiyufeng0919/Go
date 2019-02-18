package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type photo struct {
	Id string
	ImagePath string
}

//返回三项数据:photos,title,price
func detailGet(c *gin.Context){
	//假装处理商品id
	id:=c.Param(":id")
	fmt.Printf("id is:%s\n",id)

	//假装访问数据库
	//初始化photos
	buffer:=[]photo{
		photo{
			Id:"pic1",
			ImagePath:"img01.png",
		},
		photo{
			Id:"pic2",
			ImagePath:"img02.png",
		},
		photo{
			Id:"pic3",
			ImagePath:"img03.png",
		},
		photo{
			Id:"pic4",
			ImagePath:"img04.png",
		},
		photo{
			Id:"pic5",
			ImagePath:"img05.png",
		},
		photo{
			Id:"pic6",
			ImagePath:"img06.png",
		},
	}

	//初始化price
	price:=139900
	intPrice:=price / 100 //整数部分
	decPrice:=price % 100 //小数部分
	priceString:=fmt.Sprintf("%d.%02d",intPrice,decPrice)

	//gin框架返回json格式数据
	c.JSON(http.StatusOK,gin.H{
		"photos":buffer,
		"title":"开心玉凤学习Gin",
		"price":priceString,
	})
}

/*
Gin实现商品详情页
1。搭建静态html服务
2。实现ajax服务器部分
3。实现ajax html部分
*/
func main(){

	router:=gin.Default()

	//设定html文件夹路由,用于存放index.html以及css/js等文件
	router.Static("/html","./html")

	//路由，Ajax请求路由
	router.GET("/detail/:id",detailGet) //浏览器访问:http://127.0.0.1:8084/detail/123页面会打印json数据

	//服务启动
	router.Run("127.0.0.1:8084")

}
