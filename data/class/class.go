package class

import (
	"CourseManagement/data"
	"encoding/json"
	"sync"
)

type Class struct {
	ID           int
	Name         string
	StuIDList    map[int]int
	HeadmasterID int
	CourseID     int
	sync.RWMutex
}
type ClassData struct {
	tableName data.TableName
}

func NewClassData() *ClassData {
	return &ClassData{
		tableName: data.Class,
	}
}

func (d *ClassData) Add(class *Class) (int, error) {
	id, err := data.Add(d.tableName, class)
	class.ID = id
	data.ShowAllData(d.tableName)
	return class.ID, err
}

func (d *ClassData) Edit(class *Class) error {
	err := data.Edit(d.tableName, class.ID, class)
	data.ShowAllData(d.tableName)
	return err
}

func (d *ClassData) Del(id int) error {
	err := data.Del(d.tableName, id)
	data.ShowAllData(d.tableName)
	return err
}

func (d *ClassData) Get(id ...int) ([]*Class, error) {
	list := make([]*Class, 0)
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
			class, ok := v.(*Class)
			if !ok {
				continue
			}
			class.ID = i
			list = append(list, class)
		}
	} else {
		for k, v := range mp {
			class, ok := v.(*Class)
			if !ok {
				continue
			}
			class.ID = k
			list = append(list, class)
		}
	}
	return list, err
}

func (c *Class) AddStu(StuID ...int) {
	c.Lock()
	defer c.Unlock()
	if c.StuIDList == nil {
		c.StuIDList = make(map[int]int)
	}
	for _, id := range StuID {
		c.StuIDList[id] = id
	}
}

func (c *Class) DelStu(StuID ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range StuID {
		delete(c.StuIDList, id)
	}
}

func (c *Class) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}
