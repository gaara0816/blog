package controllers

import (
	"demon/blog/models"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {

	c.Data["IsArticle"] = true
	c.TplName = "article.html"
	isLogin := checkAccount(c.Ctx)
	c.Data["IsLogin"] = isLogin
	var err error
	c.Data["Articles"], err = models.ObtainAllArticles()

	if err != nil {
		beego.Error(err)
	}

}

func (c *ArticleController) Post() {
	login := checkAccount(c.Ctx)
	if !login {
		c.Redirect("/login", 302)
		return
	}
	name := c.Input().Get("name")
	content := c.Input().Get("content")
	err := models.AddArticle(name, content)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/article", 302)
}

func (c *ArticleController) Add() {
	c.TplName = "article_add.html"
}
