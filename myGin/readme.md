# gin框架实现商口详情页

## 需求

应用gin框架实现商品详情页面。定义Ajax，异步请求访问获取商品详情信息

1.创建myGin项目

**注意：该项目必须放置到$GOPATH/src目录下面**

2.创建main.go程序入口文件

主要定义如下：

（1）定义模型数据--M (模拟从数据库获取数据)

（2）定义静态页面--V (ajax请求页面)

（3）定义路由将模型数据返回页面--C (将数据渲染到页面中)

3.创建html文件夹，并将静态页面全部拷备到html文件夹中

4.修改index.html，将需要渲染的元素添加id值(用于定位)，添加ajax异步请求。

5.浏览器访问：http://127.0.0.1:8084/html/index.html即可。

##请求流程

1.浏览器访问 http://127.0.0.1:8084/html/index.html

2.页面初始化->ajax异步请求路由->后端返回结果->ajax接收结果->渲染数据

##[示例视频](https://www.imooc.com/learn/602)
