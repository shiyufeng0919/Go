package models

import "fmt"

//M层
/*
  数据库数据-》绑定到->对象(go)
  对象->持久化->数据库
*/
//假装访问了数据库
type photo struct {
	Id string
	ImagePath string
}

//返回数据:photos部分
func GetPhotos() []photo {
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
	return buffer
}

//返回数据title
func GetTitle() string{
	return "开心玉凤学习beego"
}

//返回数据，price
func GetPriceString() string{
	price:=139900
	intPrice:=price / 100 //整数部分
	decPrice:=price % 100 //小数部分

	return fmt.Sprintf("%d.%02d",intPrice,decPrice)

}