package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"starcoin-explorer-api/utils"
	"strings"
)

// Operations about object
type BaseController struct {
	beego.Controller
}

type NestPreparer interface {
	NestPreparer()
}

var esPrefix = ""

func (c *BaseController) Prepare() {
	url := c.Ctx.Input.URL()
	network := strings.Split(url, "/")[3]
	esPrefix = fmt.Sprintf("%s*", network)
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPreparer()
	}
	return
}

func (c *BaseController) Stop() {

}

func (c *BaseController) Response(result interface{}, err error, errorMessages ...*utils.ErrorMessage) {
	if err != nil {
	    utils.LogJson(err.Error())
		c.Data["json"] = map[string]interface{}{"code": 400, "message": err.Error()}
	} else if len(errorMessages) > 0 {
		errorMessage := errorMessages[0]
		c.Data["json"] = map[string]interface{}{"code": errorMessage.Code, "message": errorMessage.Message}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 200, "data": result}
	}
	c.ServeJSON()
}
