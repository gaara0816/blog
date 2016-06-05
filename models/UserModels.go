package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
)

const (
	_DB_NAME        = "data/blog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Person struct {
	Id     int64
	Name   string
	Avatar string
}

type Visitor struct {
	Person
}

type User struct {
	Person
	PhoneNum string
	Age      int8
}

type Category struct {
	Id                int64
	Title             string
	Created           time.Time `orm:"auto_now_add;type(date)"`
	Views             int64     `orm:"index"`
	ArticleTime       time.Time `orm:"auto_now_add;type(date)"`
	ArticleCount      int64
	ArticleLastUserId int64
}

type Article struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	AttachMent      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type name int8

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Person), new(Visitor), new(User), new(Category), new(Article))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := Category{Title: name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(&cate)

	if err == nil {
		return err
	}
	_, err = o.Insert(&cate)
	if err != nil {
		return err
	}

	return nil
}

func ObtainAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	qs := o.QueryTable("category")

	_, err := qs.All(&categories)
	return categories, err
}
