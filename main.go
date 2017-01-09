package main

import (
	_ "github.com/1for/inote/routers"
	"github.com/astaxie/beego"
	"github.com/1for/inote/controllers"
)

func main() {
	beego.SetLogger("file", `{"filename":"./log/inote.log"}`)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

