package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"managerStudent/myThrift/gen-go/apiservice"
	"managerStudent/serverDB/thrift/gen-go/openstars/core/bigset/generic"
)

var client *generic.TStringBigSetKVServiceClient

type ManagerStudentHandler struct {
	log map[int]*apiservice.ManagerStudent
}

func NewManagerStudentHandler() *ManagerStudentHandler {
	return &ManagerStudentHandler{log: make(map[int]*apiservice.ManagerStudent)}
}

func (p *ManagerStudentHandler) Init(ctx context.Context) (err error) {
	client.CreateStringBigSet(ctx, "SinhVien")
	client.CreateStringBigSet(ctx, "LopHocPhan")
	fmt.Println("Oke")
	return err
}

func (p *ManagerStudentHandler) PutLopHP(ctx context.Context, lopHP *apiservice.LopHocPhan) (r int32, err error) {
	if re, err := p.ExistsLopHP(ctx, lopHP.Ma); re != 1 {
		if err != nil {
			log.Fatal(err, "  myThrift/Handel.go:32")
		}
		item := generic.NewTItem()
		item.Key = []byte(lopHP.Ma)
		b, err := json.Marshal(lopHP)
		item.Value = b
		re, err := client.BsPutItem(ctx, "LopHocPhan", item)
		if err != nil {
			log.Fatal(err, "  myThrift/Handel.go:36")
		}
		if re.Ok == true {
			return 1, err
		}
		return -1, err
	} else {
		if err != nil {
			log.Fatal(err, "  myThrift/Handel.go:32")
		}
		return -2, err
	}

}

func (p *ManagerStudentHandler) AddSinhVienVaoLop(ctx context.Context, sv *apiservice.SinhVien, maLHP string) (r int32, err error) {
	if re, _ := p.ExistsSinhVienTrongLop(ctx, maLHP, sv.Ma); re != -1 {
		log.Println("Sinh vien chua ton tai trong lop")
		if check, err := p.ExistsSinhVien(ctx, sv.Ma); check != 1 {
			if err != nil {
				log.Fatal(err)
			}
			return -2, err
		}
		lhp, err := p.GetLopHocPhan(ctx, maLHP)
		if err != nil {
			log.Fatal(err)
		}
		lhp.DsSinhVien = append(lhp.DsSinhVien, sv)
		temp, err := client.BsGetItem(ctx, "LopHocPhan", []byte(maLHP))
		if err != nil {
			log.Fatal(err, " myThrift/Handel.go:54")
		}
		log.Println("----- danh sach da ton tai: ", lhp.DsSinhVien)
		temp.Item.Value, err = json.Marshal(lhp)
		if err != nil {
			log.Fatal(err, " myThrift/Handel.go:59")
		}
		re, err := client.BsPutItem(ctx, "LopHocPhan", temp.Item)
		if err != nil {
			log.Fatal(err, " myThrift/Handel.go:63")
		}
		if re.Ok == true {
			return 1, nil
		}
		return -1, err
	} else {
		log.Println("Sinh vien da ton tai trong lop")
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
	if err != nil {
		log.Fatal(err, "  myThrift/Handel.go:83")
	}
	log.Println("danh sach sinh vien hien co", ds)
	for _, item := range ds {
		if item.Ma == maSinhVien {
			return -1, err
		}
	}
	return 1, err
}

func (p *ManagerStudentHandler) GetLopHocPhan(ctx context.Context, ma string) (r *apiservice.LopHocPhan, err error) {
	re, err := client.BsGetItem(ctx, "LopHocPhan", []byte(ma))
	if err != nil {
		log.Fatal(err, " myThrift/Handel.go:75")
	}
	if re.Item == nil {
		log.Println("Lop hoc phan ko ton tai")
		return nil, err
	}
	var i *apiservice.LopHocPhan
	if err := json.Unmarshal(re.Item.Value, &i); err != nil {
		log.Fatal(err, " myThrift/Handel.go:80")
		return nil, err
	}
	return i, err
}

func (p *ManagerStudentHandler) GetLopHocPhanSlice(ctx context.Context) (r apiservice.LopHocPhanSlices, err error) {
	count, err := client.GetTotalCount(ctx, "LopHocPhan")
	if err != nil {
		log.Fatal(err, " myThrift/Handel.go:135")
	} else {
		if count > 0 {
			re, err := client.BsGetSlice(ctx, "LopHocPhan", 0, int32(count))
			if err != nil {
				log.Fatal(err)
			} else {
				var dsL = apiservice.LopHocPhanSlices{}
				var i *apiservice.LopHocPhan
				for _, item := range re.Items.Items {
					if err := json.Unmarshal(item.Value, &i); err != nil {
						log.Println(err, " myThrift/Handel.go:147")
						return nil, err
					}
					dsL = append(dsL, i)
				}
				return dsL, err
			}
		}
	}
	return nil, err
}

func (p *ManagerStudentHandler) GetSinhVienLHP(ctx context.Context, maLHP string) (r apiservice.SinhVienSlices, err error) {
	var i apiservice.LopHocPhan
	re, err := client.BsGetItem(ctx, "LopHocPhan", []byte(maLHP))
	if err != nil {
		log.Fatal(err, "  myThrift/Handel.go:123")
	}
	if re.Item == nil {
		log.Println("Khong ton tai lop")
		return nil, err
	}
	log.Println(string(re.Item.Value))
	if err := json.Unmarshal(re.Item.Value, &i); err != nil {
		return nil, err
	}
	return i.DsSinhVien, err
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
	if check, err := p.ExistsSinhVienTrongLop(ctx, maLHP, maSV); check != 1 {
		log.Println("Sinh vien ton tai can delete")
		ds, _ := p.GetSinhVienLHP(ctx, maLHP)
		var temp apiservice.SinhVienSlices
		for _, item := range ds {
			if item.Ma != maSV {
				temp = append(temp, item)
			}
		}
		lhp, err := p.GetLopHocPhan(ctx, maLHP)
		if err != nil {
			log.Fatal(err)
		}
		lhp.DsSinhVien = temp
		temp2, err := client.BsGetItem(ctx, "LopHocPhan", []byte(maLHP))
		if err != nil {
			log.Fatal(err, " myThrift/Handel.go:199")
		}
		temp2.Item.Value, err = json.Marshal(lhp)
		if err != nil {
			log.Fatal(err, " myThrift/Handel.go:203")
		}
		re, err := client.BsPutItem(ctx, "LopHocPhan", temp2.Item)
		if err != nil {
			log.Fatal(err, " myThrift/Handel.go:207")
		}
		if re.Ok == true {
			return 1, nil
		}
		return -1, err
	} else {
		return -2, err
	}
}

func (p *ManagerStudentHandler) PutSinhVien(ctx context.Context, sv *apiservice.SinhVien) (r int32, err error) {
	if r, err := p.ExistsSinhVien(ctx, sv.Ma); r != 1 {
		if err != nil {
			log.Fatal(err, "  myThrift/Handel.go:161")
		}
		log.Println("Sinh Vien chua ton tia")
		item := generic.NewTItem()
		item.Key = []byte(sv.Ma)
		b, err := json.Marshal(sv)
		log.Println(string(b))
		log.Println(err, " src/Handel.go:126")
		item.Value = b
		re, err := client.BsPutItem(ctx, "SinhVien", item)
		log.Println(err, " src/Handel.go:129")
		// re, _ := client.BsGetItem(ctx, "SinhVien", []byte(sv.Ma))
		if re.Ok == true {
			return 1, err
		}
		return -1, err
	} else {
		if err != nil {
			log.Fatal(err, "  myThrift/Handel.go:182")
		}
		return -2, err
	}
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
	if err != nil {
		log.Println(err, " myThrift/Handel.go:253")
	}
	if re.Item == nil {
		log.Println("Sinh vien ko ton tai")
		return nil, err
	}
	var i *apiservice.SinhVien
	if err := json.Unmarshal(re.Item.Value, &i); err != nil {
		log.Println(err, " myThrift/Handel.go:261")
		return nil, err
	}
	return i, err
}

func (p *ManagerStudentHandler) DelSinhVien(ctx context.Context, maSV string) (r int32, err error) {
	re, err := client.BsRemoveItem(ctx, "SinhVien", []byte(maSV))
	if err != nil {
		log.Fatal(err, "  myThrift/Handel.go:270")
	}
	if re == true {
		ds, err := p.GetLopHocPhanSlice(ctx)
		if err != nil {
			log.Fatal(err)
		}
		for _, i := range ds {
			_, err := p.PutSVOutLopHP(ctx, i.Ma, maSV)
			if err != nil {
				log.Fatal(err)
			}
		}
		return 1, err
	} else {
		return -1, err
	}
}
