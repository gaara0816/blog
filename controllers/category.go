package controllers

import (
	"demon/blog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	}
	c.Data["IsCategory"] = true
	c.TplName = "category.html"
	isLogin := checkAccount(c.Ctx)
	c.Data["IsLogin"] = isLogin
	var err error
	c.Data["Categories"], err = models.ObtainAllCategories()

	if err != nil {
		beego.Error(err)
	}

}
