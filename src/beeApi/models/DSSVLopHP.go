package models

type IDSinhVienSlices []string

type DanhSachSinhVienLopHocPhan struct {
	MaLHP string
	DsSV  IDSinhVienSlices
}

type SinhVienVaLop struct {
	MaLHP string
	MaSV  string
}
