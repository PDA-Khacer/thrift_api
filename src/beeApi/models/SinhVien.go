package models

type SinhVien struct {
	Ma       string
	HoTen    string
	GioiTinh int32
	NgaySinh string
	Sdt      string
}

type SinhVienSlices []*SinhVien
