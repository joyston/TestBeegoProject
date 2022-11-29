package routers

import (
	"minicarproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/manufacturer", &controllers.MainController{}, "*:Manufacturer")
	beego.Router("/updateInventory", &controllers.MainController{}, "*:UpdateInventory")
}
