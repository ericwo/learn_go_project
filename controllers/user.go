package controllers

import (
    "github.com/astaxie/beego"
)

type UserController struct {
    beego.Controller
}

func (this *UserController) Get() {
    this.Layout = "layout.tpl"
    this.TplName = "user/index.tpl"
}
