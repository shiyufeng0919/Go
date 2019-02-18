package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//RESTful路由，GET函数
func RESTfulGet(c *gin.Context){
	c.String(http.StatusOK,"hello gin in GET!")
}
//RESTful路由，POST函数
func RESTfulPost(c *gin.Context){
	c.String(http.StatusOK,"hello gin in POST!")
}

//提取path中的参数
func fetchId(c *gin.Context){
	id:=c.Param("id")
	c.String(http.StatusOK,fmt.Sprintf("id is:%s\n",id))
}

//组路由
func funcAction1(c *gin.Context){
	c.String(http.StatusOK,"func action1")
}
func funcAction2(c *gin.Context){
	c.String(http.StatusOK,"func action2")
}
func funcAction3(c *gin.Context){
	c.String(http.StatusOK,"func action3")
}

func main()  {
	router:=gin.Default()

	//RESTful路由
	router.GET("/RESTful",RESTfulGet) //postman->get请求->http://127.0.0.1:8082/RESTful
	router.POST("/RESTful",RESTfulPost)//postman->post请求->http://127.0.0.1:8082/RESTful

	//不支持正则路由
	//提取path中参数
	router.GET("/param/:id",fetchId) //postman->get请求->http://127.0.0.1:8082/param/gin，提取gin

	//组路由
	group1:=router.Group("/folder1")
	{
		group1.GET("/action1",funcAction1)
		group1.GET("/action2",funcAction2) //postman->get请求->http://127.0.0.1:8082/folder1/action2
		group1.GET("/action3",funcAction3)
	}
	group2:=router.Group("/folder2")
	{
		group2.GET("/action3",funcAction3) //postman->get请求->http://127.0.0.1:8082/folder2/action3
	}

	//服务启动
	router.Run("127.0.0.1:8082")
}
