package sysinit

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func init() {
	beego.LoadAppConfig("ini", "../conf/app.conf")
	beego.AddFuncMap("i18n", i18n.Tr)

	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "../conf/"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}
