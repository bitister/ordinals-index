package routers

import (
	"api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/domains/query", &controllers.Domain{}, "post:Query")
}
