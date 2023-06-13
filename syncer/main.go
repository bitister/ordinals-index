package main

import (
	"github.com/astaxie/beego"
	"syncer/ord"
)

func main() {
	syncer, err := ord.NewSyncer()
	if err != nil {
		beego.Error("err:", err.Error())
		return
	}

	syncer.Run()
}
