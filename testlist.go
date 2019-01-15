package main

import "github.com/astaxie/beego"


//通过组合的形式在自定义的Controller中将beego.Controller包含，
//这样自定义的Controller就有了Post,Get...等方法，我们只要重写
//Get方法，在Get中写入 "hello world"就可以了
type MainController struct {
	beego.Controller
}
type aomController struct {
	beego.Controller
}

func (this *MainController)Get(){
	this.Ctx.WriteString("hello world 中国人民万岁")
}
func (this *aomController)Get(){
	this.Ctx.WriteString("hello world 中国人民万22岁")
}

func main(){
	//beego.Router("/",&MainController{})
	beego.Router("/",&aomController {})
	beego.Run()
}
