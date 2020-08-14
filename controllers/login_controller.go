package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get()  {
	c.TplName = "login.html"
}

func (c *LoginController) Post()  {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username:", username, ",password:", password)

	id := models.QueryUserWithParams(username, utils.MD5(password))
	fmt.Println("id", id)

	if id > 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "登陆成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登陆失败"}
	}

	c.ServeJSON()
}
