package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"],
		beego.ControllerComments{
			Method:           "UpdateClass",
			Router:           "/",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"],
		beego.ControllerComments{
			Method:           "DeleteClassC",
			Router:           "/",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"],
		beego.ControllerComments{
			Method:           "GetClass",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"],
		beego.ControllerComments{
			Method:           "GetAllClass",
			Router:           "/all",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:ClassCController"],
		beego.ControllerComments{
			Method:           "CreateClassC",
			Router:           "/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "UpdateStudent",
			Router:           "/",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "GetStudent",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "DeleteStudent",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "AddStudentToClass",
			Router:           "/AddToClass",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "GetAllClassOfStudent",
			Router:           "/AllClass",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "GetAllStudentInClass",
			Router:           "/AllStudent",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "RemoveStudentOutClass",
			Router:           "/OutClass",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "GetAllStudent",
			Router:           "/all",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"] = append(beego.GlobalControllerRouter["managerStudent/ver2/src/apiManager/controllers:StudentController"],
		beego.ControllerComments{
			Method:           "CreateStudent",
			Router:           "/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
