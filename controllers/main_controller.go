package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]string{}
	c.ServeJSON()
}
