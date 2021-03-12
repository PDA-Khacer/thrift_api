package controllers

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"managerStudent/ver2/src/apiManager/models"
	"managerStudent/ver2/thrift/gen-go/datamanager"
)

type HeaderID struct {
	student string
	classC  string
}

var HEADER_ID = HeaderID{student: "STU_", classC: "CC_"}

type StudentController struct {
	beego.Controller
}

// get

// @Title GetStudent
// @Description create users
// @Param id	path string		true		"body for sinhvien content"
// @Success 200 models.Student
// @Failure 403 body is empty
// @router /:id [get]
func (u *StudentController) GetStudent() {
	id := u.GetString(":id")
	if id != "" {
		user, err := models.GetClient().GetStudent(context.Background(), HEADER_ID.student+id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title GetAllStudent
// @Description create users
// @Success 200 models.StudentSlice
// @Failure 403 body is empty
// @router /all [get]
func (u *StudentController) GetAllStudent() {
	user, err := models.GetClient().GetAllStudent(context.Background())
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title GetAllClassOfStudent
// @Description create users
// @Param id	path string		true		"body for sinhvien content"
// @Success 200 models.StudentSlice
// @Failure 403 body is empty
// @router /AllClass :id [get]
func (u *StudentController) GetAllClassOfStudent() {
	id := u.GetString(":id")
	if id != "" {
		user, err := models.GetClient().GetAllClassOfStudent(context.Background(), HEADER_ID.student+id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// post

// @Title CreateStudent
// @Description create users
// @Param	body		body 	models.Student	true		"body for sinhvien content"
// @Success 200 {int} 1
// @Failure 403 body is empty
// @router /create [post]
func (u *StudentController) CreateStudent() {
	var sv datamanager.Student
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	uid, err := models.GetClient().AddStudent(context.Background(), &sv)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// put

// @Title UpdateStudent
// @Description create users
// @Param	body		body 	models.StudentInfor	true		"body for sinhvien content"
// @Success 200 {int} 1
// @Failure 403 body is empty
// @router / [put]
func (u *StudentController) UpdateStudent() {
	var sv datamanager.StudentInfor
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	uid, err := models.GetClient().UpdateStudent(context.Background(), &sv)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// @Title AddStudentToClass
// @Description create users
// @Param	body		body 	models.StudentClassC	true		"body for sinhvien content"
// @Success 200 {int} 1
// @Failure 403 body is empty
// @router /AddToClass [put]
func (u *StudentController) AddStudentToClass() {
	var sv models.StudentClassC
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	uid, err := models.GetClient().AddStudentToClass(context.Background(), HEADER_ID.student+sv.StudentId, HEADER_ID.classC+sv.ClassId)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// @Title RemoveStudentOutClass
// @Description create users
// @Param	body		body 	models.StudentClassC	true		"body for sinhvien content"
// @Success 200 {int} 1
// @Failure 403 body is empty
// @router /OutClass [delete]
func (u *StudentController) RemoveStudentOutClass() {
	var sv models.StudentClassC
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	uid, err := models.GetClient().RemoveStudentInClass(context.Background(), HEADER_ID.student+sv.StudentId, HEADER_ID.classC+sv.ClassId)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// del

// @Title DeleteStudent
// @Description create users
// @Param id	path string		true		"body for sinhvien content"
// @Success 200 {int} 1
// @Failure 403 body is empty
// @router /:id [delete]
func (u *StudentController) DeleteStudent() {
	id := u.GetString(":id")
	if id != "" {
		user, err := models.GetClient().RemoveStudent(context.Background(), HEADER_ID.student+id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}
