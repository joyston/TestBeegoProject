package main

import (
	/* _ "minicarproject/routers"
	"github.com/astaxie/beego" */
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@/minicar?charset=utf8")
}

func main() {
	orm.Debug = true
	//beego.Run()
	web.Run()
}
