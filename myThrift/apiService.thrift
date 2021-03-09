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
    3: string giaoVien,
    4: SinhVienSlices dsSinhVien
}

typedef list<LopHocPhan> LopHocPhanSlices
# services

service ManagerStudent{
    # Lop Hoc Phan
    void init(),
    i32 putLopHP(1:LopHocPhan lopHP),
    i32 addSinhVienVaoLop(1:SinhVien sv, 2:string maLHP),
    i32 addSinhVienSlicesVaoLop(1:SinhVienSlices lsv, 2: string maLHP),
    i32 existsLopHP(1:string maLHP),
    i32 existsSinhVienTrongLop(1:string maLHP, 2:string maSinhVien),
    LopHocPhan getLopHocPhan(1:string ma),
    LopHocPhanSlices getLopHocPhanSlice(),
    SinhVienSlices getSinhVienLHP(1:string maLHP),
    i32 delLopHP(1:string maLHP),
    i32 putSVOutLopHP(1:string maLHP, 2:string maSV),
    # Sinh vien
    i32 putSinhVien(1:SinhVien sv),
    i32 existsSinhVien(1:string maSV),
    SinhVien getSinhVien(1:string maSV),
    i32 delSinhVien(1:string maSV)
}