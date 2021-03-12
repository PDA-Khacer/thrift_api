# struct

struct StudentInfor {
    1: string id,
    2: string name,
    3: i32 gender,
    4: string brith,
    5: string phone
}

struct ClassInfor{
    1: string id,
    2: string name,
    3: string teacher
}

struct Student{
    1: StudentInfor infor
    2: list<string> allClass
}

struct ClassC{
    1: ClassInfor infor
    2: list<string> allStudent
}

typedef list<Student> StudentSlice

typedef list<ClassC> ClassCSlice

typedef list<StudentInfor> StudentInforSlice

typedef list<ClassInfor> ClassInforSlices

# services

service ManagerStudent{
    # Post
    i32 addStudent(1:Student sv)
    i32 addClass(1:ClassC sv)
    # get
    Student getStudent(1:string id)
    ClassC getClass(1:string Id)
    StudentSlice getAllStudent()
    ClassCSlice getAllClass()
    StudentInforSlice getAllStudentInClass(1:string id)
    ClassInforSlices getAllClassOfStudent(1:string id)
    i32 isExistStudent(1:string id)
    i32 isExistClass(1:string id)
    i32 isExistStudentInClass(1:string idS, 2:string idC)
    #put
    i32 updateStudent(1:StudentInfor infor)
    i32 updateClass(1:ClassInfor infor)
    i32 addStudentToClass(1:string idS, 2:string idC)
    i32 removeStudentInClass(1:string idS, 2:string idC)
    #delete
    i32 removeStudent(1:string id)
    i32 removeClass(1:string id)
}