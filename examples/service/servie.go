package service

import (
	"github.com/houzhongjian/beer"
	"log"
)

func Rem(c *beer.Context)  {
	sess, _ := beer.Session().Start(c)
	beer.Session().Destroy(sess)
}

func Default(c *beer.Context) {
	log.Println(c.IP)
	//session,err := beer.Session().Start(c)
	//if err != nil {
	//	log.Printf("err:%+v\n",err)
	//	return
	//}
	//
	//session.Set("name","张三")
	c.Layout = "blog/layout.html"
	c.Data["title"] = "goBeer"
	c.Html("blog/index.html")
}

func Login(c *beer.Context) {
	c.Data["code"] = 1000
	c.Data["msg"] = "登录成功"
	c.Json()
}

func Detail(c *beer.Context) {
	log.Println(c.Param("id"))
	log.Println(c.Param("name"))
	log.Println(c.UserAgent)

	c.Layout = "blog/layout.html"
	c.Data["title"] = "欢迎回来"
	log.Println("userid:",c.Data["userid"])
	c.Html("blog/detail.html")
}
