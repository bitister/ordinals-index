package routers

import (
	"CryptoYes/server/api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/profit/info", &controllers.Account{}, "get:ProfitPool")
	beego.Router("/user/profit/withdraw", &controllers.Account{}, "post:WithDraw")
}
