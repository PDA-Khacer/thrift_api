package controllers

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"managerStudent/ver2/src/apiManager/models"
	"managerStudent/ver2/thrift/gen-go/datamanager"
)

type ClassCController struct {
	beego.Controller
}

// get

// @Title GetClass
// @Description create users
// @Param id	path string		true		"body for sinhvien content"
// @Success 200 models.ClassC
// @Failure 403 body is empty
// @router /:id [get]
func (u *ClassCController) GetClass() {
	id := u.GetString(":id")
	if id != "" {
		user, err := models.GetClient().GetClass(context.Background(), HEADER_ID.student+id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title GetAllClass
// @Description create users
// @Success 200 models.ClassCSlice
// @Failure 403 body is empty
// @router /all [get]
func (u *ClassCController) GetAllClass() {
	user, err := models.GetClient().GetAllClass(context.Background())
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title GetAllStudentInClass
// @Description create users
// @Param id	path string		true		"body for sinhvien content"
// @Success 200 models.ClassCSlice
// @Failure 403 body is empty
// @router /AllStudent :id [get]
func (u *StudentController) GetAllStudentInClass() {
	id := u.GetString(":id")
	if id != "" {
		user, err := models.GetClient().GetAllStudentInClass(context.Background(), HEADER_ID.classC+id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// post

// @Title CreateClassC
// @Description create users
// @Param	body		body 	models.ClassC	true		"body for sinhvien content"
// @Success 200 {int} 1
// @Failure 403 body is empty
// @router /create [post]
func (u *ClassCController) CreateClassC() {
	var sv datamanager.ClassC
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	uid, err := models.GetClient().AddClass(context.Background(), &sv)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// put
// @Title UpdateClass
// @Description create users
// @Param	body		body 	models.ClassInfor	true		"body for sinhvien content"
// @Success 200 {int} 1
// @Failure 403 body is empty
// @router / [put]
func (u *ClassCController) UpdateClass() {
	var sv datamanager.ClassInfor
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	uid, err := models.GetClient().UpdateClass(context.Background(), &sv)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// del

// @Title DeleteClassC
// @Description create users
// @Param id	path string		true		"body for sinhvien content"
// @Success 200 {int} 1
// @Failure 403 body is empty
// @router / [delete]
func (u *ClassCController) DeleteClassC() {
	id := u.GetString(":id")
	if id != "" {
		user, err := models.GetClient().RemoveClass(context.Background(), HEADER_ID.classC+id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}
