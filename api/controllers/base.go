package controllers

import (
	"encoding/json"
	"models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
)

type BaseController struct {
	beego.Controller
	controllerName string //当前控制名称
	actionName     string //当前action名称
	O              orm.Ormer
	Datetime       int64
	i18n.Locale
}

func (c *BaseController) Prepare() {
	c.Lang = "en"
	lang_head := c.Ctx.Input.Header("Accept-Language")

	if lang_head != "" && (lang_head == "zh" || lang_head == "en") {
		c.Lang = lang_head
	}
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Info("请求：", c.controllerName, c.actionName, c.Input())
	c.Datetime = time.Now().Unix()

	c.O = orm.NewOrm()
}

func (c *BaseController) Finish() {
	data, _ := json.Marshal(c.Data["json"])
	_, c.actionName = c.GetControllerAndAction()
	beego.Info("返回：", string(data))
}

//返回数据工具
func (c *BaseController) Fail(res string, data interface{}) models.Result {
	codemsgs := strings.Split(res, "|")
	codeint, _ := strconv.ParseInt(codemsgs[0], 10, 64)
	resp := models.Result{
		Code:    int(codeint),
		Status:  false,
		Message: codemsgs[1],
		Data:    data,
	}
	return resp
}

//返回数据工具
func (c *BaseController) Succ(res string, data interface{}) models.Result {
	codemsgs := strings.Split(res, "|")
	codeint, _ := strconv.ParseInt(codemsgs[0], 10, 64)
	resp := models.Result{
		Code:    int(codeint),
		Status:  true,
		Message: codemsgs[1],
		Data:    data,
	}
	return resp
}
