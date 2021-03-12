package main

import (
	"context"
	"encoding/json"
	"log"
	"managerStudent/serverDB/thrift/gen-go/openstars/core/bigset/generic"
	"managerStudent/ver2/thrift/gen-go/datamanager"
	"time"
)

var client *generic.TStringBigSetKVServiceClient

type HeaderID struct {
	student string
	classC  string
}

var HEADER_ID = HeaderID{student: "STU_", classC: "CC_"}

type TimeLog struct {
	timeAdd string
}

type ManagerStudentHandler struct {
	log map[int]*datamanager.ManagerStudent
}

func (m ManagerStudentHandler) AddStudent(ctx context.Context, sv *datamanager.Student) (r int32, err error) {
	bsKey := HEADER_ID.student + sv.Infor.ID
	if re, err := m.IsExistStudent(ctx, bsKey); err != nil {
		log.Fatal(err)
	} else if re == -1 {
		item := generic.NewTItem()
		item.Key = []byte("info")
		if j, err := json.Marshal(sv.Infor); err != nil {
			log.Fatal(err)
		} else {
			item.Value = j
		}
		if re2, err := client.BsPutItem(ctx, generic.TStringKey(bsKey), item); err != nil {
			log.Fatal(err)
		} else if re2.Ok == true {
			item2 := generic.NewTItem()
			timeAdd := TimeLog{timeAdd: time.Now().String()}
			item2.Key = []byte(bsKey)
			if j2, err := json.Marshal(timeAdd); err != nil {
				log.Fatal(err)
			} else {
				item.Value = j2
				if re3, err := client.BsPutItem(ctx, "Student", item); err != nil {
					log.Fatal(err)
				} else if re3.Ok == true {
					return 1, err
				} else {
					return -1, err
				}
			}
		} else {
			return -1, err
		}
	} else {
		return -2, err
	}
	return -99, err
}

func (m ManagerStudentHandler) AddClass(ctx context.Context, sv *datamanager.ClassC) (r int32, err error) {
	bsKey := HEADER_ID.classC + sv.Infor.ID
	if re, err := m.IsExistClass(ctx, bsKey); err != nil {
		log.Fatal(err)
	} else if re == -1 {
		item := generic.NewTItem()
		item.Key = []byte("info")
		if j, err := json.Marshal(sv.Infor); err != nil {
			log.Fatal(err)
		} else {
			item.Value = j
		}
		if re2, err := client.BsPutItem(ctx, generic.TStringKey(bsKey), item); err != nil {
			log.Fatal(err)
		} else if re2.Ok == true {
			item2 := generic.NewTItem()
			timeAdd := TimeLog{timeAdd: time.Now().String()}
			item2.Key = []byte(bsKey)
			if j2, err := json.Marshal(timeAdd); err != nil {
				log.Fatal(err)
			} else {
				item.Value = j2
				if re3, err := client.BsPutItem(ctx, "ClassC", item); err != nil {
					log.Fatal(err)
				} else if re3.Ok == true {
					return 1, err
				} else {
					return -1, err
				}
			}
		} else {
			return -1, err
		}
	} else {
		return -2, err
	}
	return -99, err
}

func (m ManagerStudentHandler) GetStudent(ctx context.Context, Id string) (r *datamanager.Student, err error) {
	if re1, err := m.IsExistStudent(ctx, Id); err != nil {
		log.Fatal(err)
	} else {
		if re1 == 1 {
			bsKey := HEADER_ID.student + Id
			var i = &datamanager.Student{}
			if re2, err := client.BsGetItem(ctx, generic.TStringKey(bsKey), []byte("info")); err != nil {
				log.Fatal(err)
			} else if re2.Item == nil {
				return nil, err
			} else if err := json.Unmarshal(re2.Item.Value, &i); err != nil {
				log.Fatal(err)
			} else {
				return i, err
			}
		} else {
			return nil, err
		}
	}
	return nil, err
}

func (m ManagerStudentHandler) GetClass(ctx context.Context, id string) (r *datamanager.ClassC, err error) {
	if re1, err := m.IsExistClass(ctx, id); err != nil {
		log.Fatal(err)
	} else {
		if re1 == 1 {
			bsKey := HEADER_ID.classC + id
			var i = &datamanager.ClassC{}
			if re2, err := client.BsGetItem(ctx, generic.TStringKey(bsKey), []byte("info")); err != nil {
				log.Fatal(err)
			} else if re2.Item == nil {
				return nil, err
			} else if err := json.Unmarshal(re2.Item.Value, &i); err != nil {
				log.Fatal(err)
			} else {
				return i, err
			}
		} else {
			return nil, err
		}
	}
	return nil, err
}

func (m ManagerStudentHandler) GetAllStudent(ctx context.Context) (r datamanager.StudentSlice, err error) {
	if count, err := client.GetTotalCount(ctx, "Student"); err != nil {
		log.Fatal(err)
	} else if count == 0 {
		return nil, err
	} else {
		var ls = datamanager.StudentSlice{}
		if lsId, err := client.BsGetSlice(ctx, "Student", 0, int32(count)); err != nil {
			log.Fatal(err)
		} else {
			for _, i := range lsId.Items.Items {
				if re2, err := m.GetStudent(ctx, string(i.Key)); err != nil {
					log.Fatal(err)
				} else {
					ls = append(ls, re2)
				}
			}
			return ls, err
		}
	}
	return nil, err
}

func (m ManagerStudentHandler) GetAllClass(ctx context.Context) (r datamanager.ClassCSlice, err error) {
	if count, err := client.GetTotalCount(ctx, "ClassC"); err != nil {
		log.Fatal(err)
	} else if count == 0 {
		return nil, err
	} else {
		var ls = datamanager.ClassCSlice{}
		if lsId, err := client.BsGetSlice(ctx, "ClassC", 0, int32(count)); err != nil {
			log.Fatal(err)
		} else {
			for _, i := range lsId.Items.Items {
				if re2, err := m.GetClass(ctx, string(i.Key)); err != nil {
					log.Fatal(err)
				} else {
					ls = append(ls, re2)
				}
			}
			return ls, err
		}
	}
	return nil, err
}

func (m ManagerStudentHandler) GetAllStudentInClass(ctx context.Context, id string) (r datamanager.StudentInforSlice, err error) {
	if count, err := client.GetTotalCount(ctx, generic.TStringKey(HEADER_ID.classC+id)); err != nil {
		log.Fatal(err)
	} else if count > 0 {
		if re, err := m.IsExistClass(ctx, id); err != nil {
			log.Fatal(err)
		} else if lsS, err := client.BsGetSlice(ctx, generic.TStringKey(HEADER_ID.classC+id), 1, int32(count)); re == 1 && err == nil {
			var result = datamanager.StudentInforSlice{}
			for _, i := range lsS.Items.Items {
				if sI, err := m.GetStudent(ctx, string(i.Key)); err != nil {
					log.Fatal(err)
				} else {
					result = append(result, sI.Infor)
				}
			}
		} else if err == nil {
			log.Fatal(err)
		}
	}
	return nil, err
}

func (m ManagerStudentHandler) GetAllClassOfStudent(ctx context.Context, id string) (r datamanager.ClassInforSlices, err error) {
	if count, err := client.GetTotalCount(ctx, generic.TStringKey(HEADER_ID.student+id)); err != nil {
		log.Fatal(err)
	} else if count > 0 {
		if re, err := m.IsExistClass(ctx, id); err != nil {
			log.Fatal(err)
		} else if lsC, err := client.BsGetSlice(ctx, generic.TStringKey(HEADER_ID.student+id), 1, int32(count)); re == 1 && err == nil {
			var result = datamanager.ClassInforSlices{}
			for _, i := range lsC.Items.Items {
				if sI, err := m.GetClass(ctx, string(i.Key)); err != nil {
					log.Fatal(err)
				} else {
					result = append(result, sI.Infor)
				}
			}
		} else if err == nil {
			log.Fatal(err)
		}
	}
	return nil, err
}

func (m ManagerStudentHandler) IsExistStudent(ctx context.Context, id string) (r int32, err error) {
	bsKey := HEADER_ID.student + id
	if re, err := client.BsExisted(ctx, generic.TStringKey(bsKey), []byte("info")); err != nil {
		log.Fatal(err)
	} else if re.Existed == true {
		return 1, err
	}
	return -1, err
}

func (m ManagerStudentHandler) IsExistClass(ctx context.Context, id string) (r int32, err error) {
	bsKey := HEADER_ID.classC + id
	if re, err := client.BsExisted(ctx, generic.TStringKey(bsKey), []byte("info")); err != nil {
		log.Fatal(err)
	} else if re.Existed == true {
		return 1, err
	}
	return -1, err
}

func (m ManagerStudentHandler) IsExistStudentInClass(ctx context.Context, idS string, idC string) (r int32, err error) {
	if re1, err := m.GetStudent(ctx, idS); err != nil {
		log.Fatal(err)
	} else if re2, err := client.BsGetItem(ctx, generic.TStringKey(HEADER_ID.student+re1.Infor.ID), generic.TItemKey(idC)); re1 != nil && err != nil {
		log.Fatal(err)
	} else if re1 != nil && re2 != nil {
		log.Println("Kiem tra studnet in class : ", re2)
		return 1, err
	}
	return -1, err
}

func (m ManagerStudentHandler) UpdateStudent(ctx context.Context, info *datamanager.StudentInfor) (r int32, err error) {
	item := generic.NewTItem()
	item.Key = []byte(HEADER_ID.student + info.ID)
	if ij, err := json.Marshal(info); err != nil {
		log.Fatal(err)
	} else {
		item.Value = ij
		if re1, err := m.IsExistStudent(ctx, HEADER_ID.student+info.ID); err != nil {
			log.Fatal(err)
		} else if re2, err := client.BsPutItem(ctx, generic.TStringKey(HEADER_ID.student+info.ID), item); re1 == 1 && err == nil && re2.Ok == true {
			return 1, err
		} else if err != nil {
			log.Fatal(err)
		} else if re2.Ok == false {
			return -1, nil
		}
	}
	return -2, err
}

func (m ManagerStudentHandler) UpdateClass(ctx context.Context, info *datamanager.ClassInfor) (r int32, err error) {
	item := generic.NewTItem()
	item.Key = []byte(HEADER_ID.classC + info.ID)
	if ij, err := json.Marshal(info); err != nil {
		log.Fatal(err)
	} else {
		item.Value = ij
		if re1, err := m.IsExistClass(ctx, HEADER_ID.classC+info.ID); err != nil {
			log.Fatal(err)
		} else if re2, err := client.BsPutItem(ctx, generic.TStringKey(HEADER_ID.classC+info.ID), item); re1 == 1 && err == nil && re2.Ok == true {
			return 1, err
		} else if err != nil {
			log.Fatal(err)
		} else if re2.Ok == false {
			return -1, nil
		}
	}
	return -2, err
}

func (m ManagerStudentHandler) AddStudentToClass(ctx context.Context, idS string, idC string) (r int32, err error) {
	item := generic.NewTItem()
	item.Key = []byte(HEADER_ID.classC + idC)
	timeAdd := TimeLog{timeAdd: time.Now().String()}
	if j2, err := json.Marshal(timeAdd); err != nil {
		log.Fatal(err)
	} else {
		item.Value = j2
	}
	if check, err := m.IsExistStudentInClass(ctx, HEADER_ID.student+idS, HEADER_ID.classC+idC); err != nil {
		log.Fatal(err)
	} else if check == -1 {
		if re1, err := client.BsPutItem(ctx, generic.TStringKey(HEADER_ID.student+idS), item); err != nil {
			log.Fatal(err)
		} else if re1.Ok == true {
			item.Key = []byte(HEADER_ID.student + idS)
			if re2, err := client.BsPutItem(ctx, generic.TStringKey(HEADER_ID.classC+idC), item); err != nil {
				log.Fatal(err)
			} else if re2.Ok == true {
				return 1, err
			}
		}
		return -1, err
	}
	return -2, err
}

func (m ManagerStudentHandler) RemoveStudentInClass(ctx context.Context, idS string, idC string) (r int32, err error) {
	if check, err := m.IsExistStudentInClass(ctx, HEADER_ID.student+idS, HEADER_ID.classC+idC); err != nil {
		log.Fatal(err)
	} else if check == 1 {
		if re1, err := client.BsRemoveItem(ctx, generic.TStringKey(HEADER_ID.student+idS), []byte(HEADER_ID.classC+idC)); err != nil {
			log.Fatal(err)
		} else if re1 == true {
			if re2, err := client.BsRemoveItem(ctx, generic.TStringKey(HEADER_ID.classC+idC), []byte(HEADER_ID.student+idS)); err != nil {
				log.Fatal(err)
			} else if re2 == true {
				return 1, err
			}
		}
		return -1, err
	}
	return -2, err
}

func (m ManagerStudentHandler) RemoveStudent(ctx context.Context, id string) (r int32, err error) {
	if re1, err := m.IsExistStudent(ctx, HEADER_ID.student+id); err != nil {
		log.Fatal(err)
	} else if re1 == 1 {
		if re2, err := client.BsRemoveItem(ctx, "Student", []byte(HEADER_ID.student+id)); err != nil {
			log.Fatal(err)
		} else if re3, err := client.RemoveAll(ctx, generic.TStringKey(HEADER_ID.student+id)); re2 == true && err == nil && re3 > 0 {
			// remove on Class
			if lsC, err := m.GetAllClass(ctx); err != nil {
				log.Fatal(err)
			} else {
				for _, i := range lsC {
					for _, i2 := range i.AllStudent {
						if re4, err := m.RemoveStudentInClass(ctx, HEADER_ID.student+id, HEADER_ID.classC+i.Infor.ID); i2 == id && re4 == -1 {
							return -1, err
						} else if err != nil {
							log.Fatal(err)
						} else if i2 == id && re4 == 1 {
							break
						}
					}
				}
			}
			return 1, err
		} else if err != nil {
			log.Fatal(err)
		} else {
			return -1, err
		}
	}
	return -2, err
}

func (m ManagerStudentHandler) RemoveClass(ctx context.Context, id string) (r int32, err error) {
	if re1, err := m.IsExistClass(ctx, HEADER_ID.classC+id); err != nil {
		log.Fatal(err)
	} else if re1 == 1 {
		if re2, err := client.BsRemoveItem(ctx, "ClassC", []byte(HEADER_ID.classC+id)); err != nil {
			log.Fatal(err)
		} else if re3, err := client.RemoveAll(ctx, generic.TStringKey(HEADER_ID.classC+id)); re2 == true && err == nil && re3 > 0 {
			// remove on student
			if lsS, err := m.GetAllStudent(ctx); err != nil {
				log.Fatal(err)
			} else {
				for _, i := range lsS {
					for _, i2 := range i.AllClass {
						if re4, err := m.RemoveStudentInClass(ctx, HEADER_ID.student+i.Infor.ID, HEADER_ID.classC+id); i2 == id && re4 == -1 {
							return -1, err
						} else if err != nil {
							log.Fatal(err)
						} else if i2 == id && re4 == 1 {
							break
						}
					}
				}
			}
			return 1, err
		} else if err != nil {
			log.Fatal(err)
		} else {
			return -1, err
		}
	}
	return -2, err
}

func NewManagerStudentHandler() *ManagerStudentHandler {
	return &ManagerStudentHandler{log: make(map[int]*datamanager.ManagerStudent)}
}
