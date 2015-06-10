package routers

import (
	"github.com/duguying/jikexueyuan/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
