package teacher

import (
	"CourseManagement/common"
	"CourseManagement/data"
	"encoding/json"
	"time"
)

type Teacher struct {
	ID       int
	Name     string
	Gender   common.Gender
	Birthday time.Time
}
type TeacherData struct {
	tableName data.TableName
}

func NewTeacherData() *TeacherData {
	return &TeacherData{
		tableName: data.Teacher,
	}
}

func (d *TeacherData) Add(teacher *Teacher) (int, error) {
	id, err := data.Add(d.tableName, teacher)
	teacher.ID = id
	data.ShowAllData(d.tableName)
	return teacher.ID, err
}

func (d *TeacherData) Edit(teacher *Teacher) error {
	err := data.Edit(d.tableName, teacher.ID, teacher)
	data.ShowAllData(d.tableName)
	return err
}

func (d *TeacherData) Del(id int) error {
	err := data.Del(d.tableName, id)
	data.ShowAllData(d.tableName)
	return err
}

func (d *TeacherData) Get(id ...int) ([]*Teacher, error) {
	list := make([]*Teacher, 0)
	mp, err := data.Get(d.tableName, id...)
	if err != nil {
		return nil, err
	}
	if len(id) > 0 {
		for _, i := range id {
			v, ok := mp[i]
			if !ok {
				continue
			}
			teacher, ok := v.(*Teacher)
			if !ok {
				continue
			}
			teacher.ID = i
			list = append(list, teacher)
		}
	} else {
		for k, v := range mp {
			teacher, ok := v.(*Teacher)
			if !ok {
				continue
			}
			teacher.ID = k
			list = append(list, teacher)
		}
	}
	return list, err
}

func (t *Teacher) String() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}
