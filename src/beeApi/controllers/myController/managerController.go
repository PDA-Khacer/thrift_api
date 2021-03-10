package myController

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"managerStudent/myThrift/gen-go/apiservice"
	"managerStudent/src/beeApi/models"
)

// Operations about Users
type ManagerController struct {
	beego.Controller
}

var defaultCtx = context.Background()

/*
{
"ma": "01",
"hoTen": "Anh",
"gioiTinh": 0,
"ngaySinh": "12/12/1999",
"sdt" : "1231231"
}
*/

// @Title CreateStudent
// @Description create users
// @Param	body		body 	models.SinhVien	true		"body for sinhvien content"
// @Success 200 {string} models.SinhVien[Ma]
// @Failure 403 body is empty
// @router /student/create [post]
func (u *ManagerController) CreateStudent() {
	var sv apiservice.SinhVien
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	uid, err := models.GetClient().PutSinhVien(defaultCtx, &sv)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// @Title CreateClass
// @Description create Class
// @Param	body		body 	models.LopHocPhan	true		"body for Class content"
// @Success 200 {string} models.LopHocPhan[Ma]
// @Failure 403 body is empty
// @router /class/create [post]
func (u *ManagerController) CreateClass() {
	var lhp apiservice.LopHocPhan
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &lhp)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	uid, err := models.GetClient().PutLopHP(defaultCtx, &lhp)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// @Title AddStudentToClass
// @Description Add a student to class
// @Param	sv		body  models.SinhVien	true		"body for Class content"
// @Param	ma_lop		path 	string	true		"The key for staticblock"
// @Success 200 {string} result
// @Failure 403 body is empty
// @router /class/add-student/:ma_lop [put]
func (u *ManagerController) AddStudentToClass() {
	var sv apiservice.SinhVien
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	idC := u.GetString(":ma_lop")
	log.Println("ma lop ", idC)
	uid, err := models.GetClient().AddSinhVienVaoLop(defaultCtx, &sv, idC)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// @Title AddStudentsToClass
// @Description Add students to class
// @Param	sv		body  models.SinhVienSlices	true		"body for Class content"
// @Param	maLop		path 	string	true		"The key for staticblock"
// @Success 200 {string} result
// @Failure 403 body is empty
// @router /class/add-students/ [put]
func (u *ManagerController) AddStudentsToClass() {
	var sv apiservice.SinhVienSlices
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	if err != nil {
		log.Fatal("Error Convert To Json")
	}
	idC := u.GetString(":sv")
	uid, err := models.GetClient().AddSinhVienSlicesVaoLop(defaultCtx, sv, idC)
	if err != nil {
		log.Println(err)
	}
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// @Title GetStudent
// @Description get user by uid
// @Param	ma		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SinhVien
// @Failure 403 :ma is empty
// @router /student/:ma [get]
func (u *ManagerController) GetStudent() {
	uid := u.GetString(":ma")
	if uid != "" {
		user, err := models.GetClient().GetSinhVien(defaultCtx, uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title GetClass
// @Description get user by uid
// @Param	ma		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.LopHocPhan
// @Failure 403 :uid is empty
// @router /class/:ma [get]
func (u *ManagerController) GetClass() {
	uid := u.GetString(":ma")
	if uid != "" {
		user, err := models.GetClient().GetLopHocPhan(defaultCtx, uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title GetStudentInClass
// @Description get Student by uid class
// @Param	ma		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SinhVienSlices
// @Failure 403 :uid is empty
// @router /class/get-student/:ma [get]
func (u *ManagerController) GetStudentInClass() {
	uid := u.GetString(":ma")
	if uid != "" {
		user, err := models.GetClient().GetSinhVienLHP(defaultCtx, uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title DeleteStudent
// @Description create users
// @Param	ma		path 	string	true		"The key for staticblock"
// @Success 200 {int} models.SinhVien[Ma]
// @Failure 403 body is empty
// @router /student/delete/:ma [delete]
func (u *ManagerController) DeleteStudent() {
	uid := u.GetString(":ma")
	if uid != "" {
		user, err := models.GetClient().DelSinhVien(defaultCtx, uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title DeleteClass
// @Description create users
// @Param	ma		path 	string	true		"The key for staticblock"
// @Success 200 {string} models.LopHocPhan[Ma]
// @Failure 403 body is empty
// @router /class/delete/:ma [delete]
func (u *ManagerController) DeleteClass() {
	uid := u.GetString(":ma")
	if uid != "" {
		user, err := models.GetClient().DelLopHP(defaultCtx, uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title SearchStudent
// @Description create users
// @Param	key		path 	string	true		"The key for staticblock"
// @Success 200 {string} models.SinhVienSlices
// @Failure 403 body is empty
// @router /class/search/:key [get]
func (u *ManagerController) SearchStudent() {
	uid := u.GetString(":key")
	if uid != "" {
		user, err := models.GetClient().SearchSinhVien(defaultCtx, uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}
