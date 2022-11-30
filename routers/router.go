package routers

import (
	"test/minicar/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/manufacturer", &controllers.MainController{}, "*:Manufacturer")
	beego.Router("/updateInventory", &controllers.MainController{}, "*:UpdateInventory")
}
