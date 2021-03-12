package models

type ClassInfor struct {
	ID      string
	Name    string
	Teacher string
}

type ClassC struct {
	Infor      *ClassInfor
	AllStudent []string
}

type ClassCSlice []*ClassC
