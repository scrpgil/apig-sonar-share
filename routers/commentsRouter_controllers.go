package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/xiaca/apig-sonar-share/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/xiaca/apig-sonar-share/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/coins`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
