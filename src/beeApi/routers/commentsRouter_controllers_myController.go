package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
        beego.ControllerComments{
            Method: "CreateStudent",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
