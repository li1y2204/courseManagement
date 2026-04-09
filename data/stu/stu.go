package stu

import (
	"CourseManagement/common"
	"CourseManagement/data"
	"encoding/json"
	"sync"
	"time"
)

type Stu struct {
	ID           int
	Name         string
	Gender       common.Gender
	Birthday     time.Time
	CourseIDList map[int]int
	ClassIDList  map[int]int
	sync.RWMutex
}

type StuData struct {
	tableName data.TableName
}

func NewStuData() *StuData { ///?
	return &StuData{
		tableName: data.Stu,
	}
}

func (d *StuData) Add(stu *Stu) (int, error) {
	id, err := data.Add(d.tableName, stu)
	stu.ID = id
	data.ShowAllData(d.tableName) ///
	return stu.ID, err
}

func (d *StuData) Edit(stu *Stu) error {
	err := data.Edit(d.tableName, stu.ID, stu)
	data.ShowAllData(d.tableName)
	return err
}

func (d *StuData) Del(id int) error {
	err := data.Del(d.tableName, id)
	data.ShowAllData(d.tableName)
	return err
}

func (d *StuData) Get(id ...int) ([]*Stu, error) {
	list := make([]*Stu, 0)
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
			stu, ok := v.(*Stu) ///
			if !ok {
				continue
			}
			stu.ID = i
			list = append(list, stu)
		}
	} else {
		for k, v := range mp {
			stu, ok := v.(*Stu)
			if !ok {
				continue
			}
			stu.ID = k
			list = append(list, stu)
		}
	}
	return list, nil
}

func (c *Stu) AddClass(ClassID ...int) {
	c.Lock()
	defer c.Unlock()
	if c.ClassIDList == nil {
		c.ClassIDList = make(map[int]int)
	}
	for _, id := range ClassID {
		c.ClassIDList[id] = id
	}
}

func (c *Stu) DelClass(ClassID ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range ClassID {
		delete(c.ClassIDList, id)
	}
}

func (c *Stu) AddCourse(CourseID ...int) {
	c.Lock()
	defer c.Unlock()
	if c.CourseIDList == nil {
		c.CourseIDList = make(map[int]int)
	}
	for _, id := range CourseID {
		c.CourseIDList[id] = id
	}
}

func (c *Stu) DelCourse(CourseID ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range CourseID {
		delete(c.CourseIDList, id)
	}
}

func (s *Stu) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}
