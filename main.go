package main

import (
	"demon/blog/controllers"
	"demon/blog/models"
	_ "demon/blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.AutoRouter(&controllers.ArticleController{})
	beego.Run()
}
