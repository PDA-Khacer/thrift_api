package models

type TimeLog struct {
	timeAdd string
}

type StudentInfor struct {
	ID     string
	Name   string
	Gender int32
	Brith  string
	Phone  string
}

type Student struct {
	Infor    *StudentInfor
	AllClass []string
}

type StudentSlice []*Student

type StudentClassC struct {
	StudentId string
	ClassId   string
}
