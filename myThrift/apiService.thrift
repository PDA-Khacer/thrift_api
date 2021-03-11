# struct

struct SinhVien {
    1: string ma,
    2: string hoTen,
    3: i32 gioiTinh,
    4: string ngaySinh,
    5: string sdt
}

typedef list<SinhVien> SinhVienSlices

struct LopHocPhan{
    1: string ma,
    2: string ten,
    3: string giaoVien
}

typedef list<LopHocPhan> LopHocPhanSlices

typedef list<string> IDSinhVienSlices

struct DanhSachSinhVienLopHocPhan{
    1:string maLHP,
    2:IDSinhVienSlices dsSV
}

# services

service ManagerStudent{
    # Lop Hoc Phan
    void init(),
    i32 addLopHP(1:LopHocPhan lopHP),
    i32 addSinhVienVaoLop(1:string sv, 2:string maLHP),
    i32 addSinhVienSlicesVaoLop(1:SinhVienSlices lsv, 2: string maLHP),
    i32 existsLopHP(1:string maLHP),
    i32 existsSinhVienTrongLop(1:string maLHP, 2:string maSinhVien),
    LopHocPhan getLopHocPhan(1:string ma),
    LopHocPhanSlices getLopHocPhanSlice(),
    LopHocPhanSlices getLopHocPhanOfSinhVien(1:string ma),
    SinhVienSlices getSinhVienLHP(1:string maLHP),
    i32 delLopHP(1:string maLHP),
    i32 putSVOutLopHP(1:string maLHP, 2:string maSV),
    # Sinh vien
    i32 addSinhVien(1:SinhVien sv),
    i32 existsSinhVien(1:string maSV),
    SinhVien getSinhVien(1:string maSV),
    i32 delSinhVien(1:string maSV),
    SinhVienSlices searchSinhVien(1:string key),
    #
    i32 UpdateSinhVien(1:SinhVien sv),
    i32 UpdateLopHP(1:LopHocPhan lhp)
}