package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "GetClass",
			Router:           "/class/:ma",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "AddStudentToClass",
			Router:           "/class/add-student/:ma_lop",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "AddStudentsToClass",
			Router:           "/class/add-students/",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "CreateClass",
			Router:           "/class/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "DeleteClass",
			Router:           "/class/delete/:ma",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "GetStudentInClass",
			Router:           "/class/get-student/:ma",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "SearchStudent",
			Router:           "/class/search/:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "GetStudent",
			Router:           "/student/:ma",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "CreateStudent",
			Router:           "/student/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"] = append(beego.GlobalControllerRouter["managerStudent/src/beeApi/controllers/myController:ManagerController"],
		beego.ControllerComments{
			Method:           "DeleteStudent",
			Router:           "/student/delete/:ma",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
