package main

import (
	"starcoin-api/db"
	_ "starcoin-api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	db.ConnectElasticsearch()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
