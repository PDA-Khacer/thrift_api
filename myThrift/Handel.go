package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"managerStudent/myThrift/gen-go/apiservice"
	"managerStudent/serverDB/thrift/gen-go/openstars/core/bigset/generic"
	"sync"
)

var doOnce sync.Once
var client *generic.TStringBigSetKVServiceClient

type ManagerStudentHandler struct {
	log map[int]*apiservice.ManagerStudent
}

func NewManagerStudentHandler() *ManagerStudentHandler {
	doOnce.Do(func() {
		fmt.Println("Run once - first time, loading...")
		client = Connect()
	})
	fmt.Println("Run this every time")
	return &ManagerStudentHandler{log: make(map[int]*apiservice.ManagerStudent)}
}

func (p *ManagerStudentHandler) Init(ctx context.Context) (err error) {
	client.CreateStringBigSet(ctx, "SinhVien")
	client.CreateStringBigSet(ctx, "LopHocPhan")
	fmt.Println("Oke")
	return err
}

func (p *ManagerStudentHandler) PutLopHP(ctx context.Context, lopHP *apiservice.LopHocPhan) (r int32, err error) {
	item := generic.NewTItem()
	item.Key = []byte(lopHP.Ma)
	b, err := json.Marshal(lopHP)
	item.Value = []byte(b)
	client.BsPutItem(ctx, "LopHocPhan", item)
	return 0, err
}

func (p *ManagerStudentHandler) AddSinhVienVaoLop(ctx context.Context, sv *apiservice.SinhVien, maLHP string) (r int32, err error) {
	if re, _ := p.ExistsSinhVienTrongLop(ctx, maLHP, maLHP); re != -1 {
		ds, _ := p.GetSinhVienLHP(ctx, maLHP)
		ds = append(ds, sv)
		return 0, err
	} else {
		fmt.Println("Sinh vien da ton tai")
		return -1, err
	}
}

func (p *ManagerStudentHandler) AddSinhVienSlicesVaoLop(ctx context.Context, lsv apiservice.SinhVienSlices, maLHP string) (r int32, err error) {
	panic("implement me")
}

func (p *ManagerStudentHandler) ExistsLopHP(ctx context.Context, maLHP string) (r int32, err error) {

	re, _ := client.BsExisted(ctx, "LopHocPhan", []byte(maLHP))
	if re.Existed == true {
		return 1, err
	}
	return 0, err
}

func (p *ManagerStudentHandler) ExistsSinhVienTrongLop(ctx context.Context, maLHP string, maSinhVien string) (r int32, err error) {
	ds, err := p.GetSinhVienLHP(ctx, maLHP)
	for _, item := range ds {
		if item.Ma == maSinhVien {
			return -1, err
		}
	}
	return 1, err
}

func (p *ManagerStudentHandler) GetLopHocPhan(ctx context.Context, ma string) (r *apiservice.LopHocPhan, err error) {
	re, err := client.BsGetItem(ctx, "LopHocPhan", []byte(ma))
	var i *apiservice.LopHocPhan
	if err := json.Unmarshal([]byte(re.Item.Value), i); err != nil {
		return nil, err
	}
	return i, err
}

func (p *ManagerStudentHandler) GetLopHocPhanSlice(ctx context.Context) (r apiservice.LopHocPhanSlices, err error) {
	panic("implement me")
}

func (p *ManagerStudentHandler) GetSinhVienLHP(ctx context.Context, maLHP string) (r apiservice.SinhVienSlices, err error) {
	var i apiservice.LopHocPhan
	re, _ := client.BsGetItem(ctx, "LopHocPhan", []byte(maLHP))
	value := re.Item.Value
	if err := json.Unmarshal([]byte(value), &i); err != nil {
		return i.DsSinhVien, err
	}
	return nil, err
}

func (p *ManagerStudentHandler) DelLopHP(ctx context.Context, maLHP string) (r int32, err error) {
	re, err := client.BsRemoveItem(ctx, "LopHocPhan", []byte(maLHP))
	if re == true {
		return 1, err
	} else {
		return -1, err
	}
}

func (p *ManagerStudentHandler) PutSVOutLopHP(ctx context.Context, maLHP string, maSV string) (r int32, err error) {
	ds, _ := p.GetSinhVienLHP(ctx, maLHP)
	var temp apiservice.SinhVienSlices
	for _, item := range ds {
		if item.Ma != maSV {
			temp = append(temp, item)
		}
	}
	lhp, _ := p.GetLopHocPhan(ctx, maLHP)
	lhp.DsSinhVien = temp
	return 1, err
}

func (p *ManagerStudentHandler) PutSinhVien(ctx context.Context, sv *apiservice.SinhVien) (r int32, err error) {
	item := generic.NewTItem()
	item.Key = []byte(sv.Ma)
	b, err := json.Marshal(sv)
	log.Println(err, " src/handel.go:126")
	item.Value = []byte(b)
	re, err := client.BsPutItem(ctx, "SinhVien", item)
	log.Println(err, " src/handel.go:129")
	// re, _ := client.BsGetItem(ctx, "SinhVien", []byte(sv.Ma))
	if re.Ok == true {
		return 1, err
	}
	return -1, err
}

func (p *ManagerStudentHandler) ExistsSinhVien(ctx context.Context, maSV string) (r int32, err error) {
	re, _ := client.BsExisted(ctx, "SinhVien", []byte(maSV))
	if re.Existed == true {
		return 1, err
	}
	return 0, err
}

func (p *ManagerStudentHandler) GetSinhVien(ctx context.Context, maSV string) (r *apiservice.SinhVien, err error) {
	re, err := client.BsGetItem(ctx, "SinhVien", []byte(maSV))
	var i *apiservice.SinhVien
	if err := json.Unmarshal([]byte(re.Item.Value), i); err != nil {
		return nil, err
	}
	return i, err
}

func (p *ManagerStudentHandler) DelSinhVien(ctx context.Context, maSV string) (r int32, err error) {
	re, err := client.BsRemoveItem(ctx, "SinhVien", []byte(maSV))
	if re == true {
		return 1, err
	} else {
		return -1, err
	}
}
