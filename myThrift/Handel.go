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

func (p *ManagerStudentHandler) AddLopHP(ctx context.Context, lopHP *apiservice.LopHocPhan) (r int32, err error) {
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
		item2 := generic.NewTItem()
		item2.Key = []byte(lopHP.Ma)
		var b2 = &apiservice.DanhSachSinhVienLopHocPhan{MaLHP: lopHP.Ma, DsSV: []string{}}
		item2.Value, err = json.Marshal(b2)
		if err != nil {
			log.Fatal(err)
		}
		re1, err := client.BsPutItem(ctx, "DanhSachSinhVienLopHocPhan", item2)
		if err != nil {
			log.Fatal(err)
		}
		if re1.Ok == false {
			return -3, err
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

func (p *ManagerStudentHandler) GetLopHocPhanOfSinhVien(ctx context.Context, ma string) (r apiservice.LopHocPhanSlices, err error) {
	if r, e := p.ExistsSinhVien(ctx, ma); e != nil {
		log.Fatal(e)
	} else {
		if r == -1 {
			return nil, err
		} else {
			count, e := client.GetTotalCount(ctx, "DanhSachSinhVienLopHocPhan")
			if e != nil {
				log.Fatal(e)
			}
			dsSVLHP, e := client.BsGetSlice(ctx, "DanhSachSinhVienLopHocPhan", 0, int32(count))
			if e != nil {
				log.Fatal(e)
			}
			var result = apiservice.LopHocPhanSlices{}
			for _, i := range dsSVLHP.Items.Items {
				if r2, e := p.ExistsSinhVienTrongLop(ctx, string(i.Key), ma); e != nil {
					log.Fatal(e)
				} else {
					if r2 == 1 {
						if lhp, e := p.GetLopHocPhan(ctx, string(i.Key)); e != nil {
							log.Fatal(e)
						} else {
							if lhp != nil {
								result = append(result, lhp)
							}
						}
					}
				}
			}
			return result, err
		}
	}
	return nil, err
}

func (p *ManagerStudentHandler) AddSinhVien(ctx context.Context, sv *apiservice.SinhVien) (r int32, err error) {
	if r, err := p.ExistsSinhVien(ctx, sv.Ma); r != 1 {
		if err != nil {
			log.Fatal(err, "  myThrift/Handel.go:161")
		}
		log.Println("Sinh Vien chua ton tia")
		item := generic.NewTItem()
		item.Key = []byte(sv.Ma)
		b, err := json.Marshal(sv)
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

func (p *ManagerStudentHandler) UpdateSinhVien(ctx context.Context, sv *apiservice.SinhVien) (r int32, err error) {
	if re, err := p.ExistsSinhVien(ctx, sv.Ma); err != nil {
		log.Fatal(err)
	} else {
		if re != -1 {
			item := generic.NewTItem()
			item.Key = []byte(sv.Ma)
			b, err := json.Marshal(sv)
			if err != nil {
				log.Fatal(err)
			}
			item.Value = b
			re, err := client.BsPutItem(ctx, "SinhVien", item)
			if err != nil {
				log.Fatal(err)
			}
			if re.Ok == true {
				return 1, err
			}
			return -1, err
		} else {
			log.Println("Sinh vien khong ton tai")
			return -2, nil
		}
	}
	return 0, err
}

func (p *ManagerStudentHandler) UpdateLopHP(ctx context.Context, lhp *apiservice.LopHocPhan) (r int32, err error) {
	if re, err := p.ExistsLopHP(ctx, lhp.Ma); err != nil {
		log.Fatal(err)
	} else {
		if re != -1 {
			item := generic.NewTItem()
			item.Key = []byte(lhp.Ma)
			b, err := json.Marshal(lhp)
			if err != nil {
				log.Fatal(err)
			}
			item.Value = b
			re, err := client.BsPutItem(ctx, "LopHocPhan", item)
			if err != nil {
				log.Fatal(err)
			}
			if re.Ok == true {
				return 1, err
			}
			return -1, err
		} else {
			log.Println("Lop hoc phan khong ton tai")
			return -2, nil
		}
	}
	return 0, err
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

func (p *ManagerStudentHandler) AddSinhVienVaoLop(ctx context.Context, sv string, maLHP string) (r int32, err error) {
	if check, err := p.ExistsSinhVien(ctx, sv); check != 1 {
		if err != nil {
			log.Fatal(err)
		}
		return -2, err
	} else {
		if re, _ := p.ExistsSinhVienTrongLop(ctx, maLHP, sv); re != -1 {
			log.Println("Sinh vien chua ton tai trong lop")
			temp, err := client.BsGetItem(ctx, "DanhSachSinhVienLopHocPhan", []byte(maLHP))
			if err != nil {
				log.Fatal(err, " myThrift/Handel.go:54")
			}
			var dssv = &apiservice.DanhSachSinhVienLopHocPhan{}
			if err := json.Unmarshal(temp.Item.Value, dssv); err != nil {
				log.Fatal(err)
			} else {
				dssv.DsSV = append(dssv.DsSV, sv)
			}
			b, err := json.Marshal(dssv)
			if err != nil {
				log.Fatal(err)
			}
			temp.Item.Value = b
			token, err := client.BsPutItem(ctx, "DanhSachSinhVienLopHocPhan", temp.Item)
			if err != nil {
				log.Fatal(err)
			}
			if token.Ok == true {
				return 1, nil
			}
		} else {
			log.Println("Sinh vien da ton tai trong lop")
			return -1, err
		}
	}
	return -1, err
}

func (p *ManagerStudentHandler) AddSinhVienSlicesVaoLop(ctx context.Context, lsv apiservice.SinhVienSlices, maLHP string) (r int32, err error) {
	panic("implement me")
}

func (p *ManagerStudentHandler) ExistsLopHP(ctx context.Context, maLHP string) (r int32, err error) {

	re, _ := client.BsExisted(ctx, "LopHocPhan", []byte(maLHP))
	if re.Existed == true {
		return 1, err
	}
	return -1, err
}

func (p *ManagerStudentHandler) ExistsSinhVienTrongLop(ctx context.Context, maLHP string, maSinhVien string) (r int32, err error) {
	var i apiservice.DanhSachSinhVienLopHocPhan
	re, err := client.BsGetItem(ctx, "DanhSachSinhVienLopHocPhan", []byte(maLHP))
	if err != nil {
		log.Fatal(err, "  myThrift/Handel.go:123")
	}
	if re.Item == nil {
		log.Println("Khong ton tai lop")
		return -2, err
	}
	log.Println(string(re.Item.Value))
	if err := json.Unmarshal(re.Item.Value, &i); err != nil {
		return -1, err
	}
	for _, item := range i.DsSV {
		if item == maSinhVien {
			return 1, err
		}
	}
	return -1, err
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
	var i apiservice.DanhSachSinhVienLopHocPhan
	re, err := client.BsGetItem(ctx, "DanhSachSinhVienLopHocPhan", []byte(maLHP))
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
	var result = apiservice.SinhVienSlices{}
	for _, item := range i.DsSV {
		var sv = &apiservice.SinhVien{}
		sv, err = p.GetSinhVien(ctx, item)
		if err != nil {
			log.Fatal()
		}
		result = append(result, sv)
	}
	return result, err
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
		ds, err := client.BsGetItem(ctx, "DanhSachSinhVienLopHocPhan", []byte(maLHP))
		if err != nil {
			log.Fatal(err)
		}
		var temp = &apiservice.DanhSachSinhVienLopHocPhan{}
		var tempS = apiservice.IDSinhVienSlices{}
		if err = json.Unmarshal(ds.Item.Value, &temp); err != nil {
			log.Fatal(err)
		} else {
			for _, item := range temp.DsSV {
				if item != maSV {
					tempS = append(tempS, item)
				}
			}
			ds.Item.Value, err = json.Marshal(tempS)
			if err != nil {
				log.Fatal(err)
			}
			token, err := client.BsPutItem(ctx, "DanhSachSinhVienLopHocPhan", ds.Item)
			if err != nil {
				log.Fatal(err)
			}
			if token.Ok == true {
				return 1, nil
			}
		}
		return -1, err
	} else {
		return -2, err
	}
}

func (p *ManagerStudentHandler) ExistsSinhVien(ctx context.Context, maSV string) (r int32, err error) {
	re, _ := client.BsExisted(ctx, "SinhVien", []byte(maSV))
	if re.Existed == true {
		return 1, err
	}
	return -1, err
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

func (p *ManagerStudentHandler) SearchSinhVien(ctx context.Context, key string) (r apiservice.SinhVienSlices, err error) {
	keyEnd := []byte(key)
	keyEnd[len(keyEnd)-1] += 1
	log.Println("key: ", []byte(key))
	log.Println("keyEnd : ", keyEnd)
	re, err := client.BsRangeQuery(ctx, "SinhVien", []byte(key), keyEnd)
	if err != nil {
		log.Fatal(err)
	}
	var dsv = apiservice.SinhVienSlices{}
	for _, i := range re.Items.Items {
		var sv *apiservice.SinhVien
		if err := json.Unmarshal(i.Value, &sv); err != nil {
			log.Fatal(err)
		}
		dsv = append(dsv, sv)
	}
	return dsv, err
}
