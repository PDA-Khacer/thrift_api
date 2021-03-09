package myController

import (
	"context"
	"encoding/json"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/astaxie/beego"
	"log"
	"managerStudent/myThrift/gen-go/apiservice"
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
// @Param	body		body 	apiservice.SinhVien	true		"body for sinhvien content"
// @Success 200 {int} apiservice.SinhVien.Ma
// @Failure 403 body is empty
// @router / [post]
func (u *ManagerController) CreateStudent() {
	var sv apiservice.SinhVien
	json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
	// ----------------------------
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("127.0.0.1:7777")
	if err != nil {
		log.Fatal("Error opening socket:", err)
	}
	transportFactory := thrift.NewTTransportFactory()
	transport, err = transportFactory.GetTransport(transport)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	if err != nil {
		log.Fatal(err)
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := apiservice.NewManagerStudentClient(thrift.NewTStandardClient(iprot, oprot))
	// ----------------------------
	uid, _ := client.PutSinhVien(defaultCtx, &sv)
	u.Data["json"] = map[string]int32{"State": uid}
	u.ServeJSON()
}

// @Title GetStudent
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} apiservice.SinhVien
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *ManagerController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		// ----------------------------
		var transport thrift.TTransport
		var err error
		transport, err = thrift.NewTSocket("127.0.0.1:7777")
		if err != nil {
			log.Fatal("Error opening socket:", err)
		}
		transportFactory := thrift.NewTTransportFactory()
		transport, err = transportFactory.GetTransport(transport)
		protocolFactory := thrift.NewTCompactProtocolFactory()
		if err != nil {
			log.Fatal(err)
		}
		iprot := protocolFactory.GetProtocol(transport)
		oprot := protocolFactory.GetProtocol(transport)
		client := apiservice.NewManagerStudentClient(thrift.NewTStandardClient(iprot, oprot))
		// ----------------------------
		user, err := client.GetSinhVien(defaultCtx, uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

