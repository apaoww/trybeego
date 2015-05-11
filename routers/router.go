package routers

import (
	"github.com/apaoww/trybeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
