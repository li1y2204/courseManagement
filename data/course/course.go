package course

import (
	"CourseManagement/data"
	"encoding/json"
	"sync"
)

type Course struct {
	ID            int
	Name          string
	TeacherIDList map[int]int
	StuIDList     map[int]int
	ClassIDList   map[int]int
	LastClassID   int
	sync.RWMutex
}

type CourseData struct {
	tableName data.TableName
}

func NewCourseData() *CourseData { ///?
	return &CourseData{
		tableName: data.Course,
	}
}

func (d *CourseData) Add(course *Course) (int, error) {
	id, err := data.Add(d.tableName, course)
	course.ID = id
	data.ShowAllData(d.tableName) ///
	return course.ID, err
}

func (d *CourseData) Edit(course *Course) error {
	err := data.Edit(d.tableName, course.ID, course)
	data.ShowAllData(d.tableName)
	return err
}

func (d *CourseData) Del(id int) error {
	err := data.Del(d.tableName, id)
	data.ShowAllData(d.tableName)
	return err
}

func (d *CourseData) Get(id ...int) ([]*Course, error) {
	list := make([]*Course, 0)
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
			course, ok := v.(*Course) ///
			if !ok {
				continue
			}
			course.ID = i
			list = append(list, course)
		}
	} else {
		for k, v := range mp {
			course, ok := v.(*Course)
			if !ok {
				continue
			}
			course.ID = k
			list = append(list, course)
		}
	}
	return list, err
}

func (c *Course) AddTeacher(teacherID ...int) {
	c.Lock()
	defer c.Unlock()
	if c.TeacherIDList == nil {
		c.TeacherIDList = make(map[int]int)
	}
	for _, id := range teacherID {
		c.TeacherIDList[id] = id
	}
}

func (c *Course) AddStu(StuID ...int) {
	c.Lock()
	defer c.Unlock()
	if c.StuIDList == nil {
		c.StuIDList = make(map[int]int)
	}
	for _, id := range StuID {
		c.StuIDList[id] = id
	}
}

func (c *Course) DelStu(StuID ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range StuID {
		delete(c.StuIDList, id)
	}
}

func (c *Course) AddClass(ClassID ...int) {
	c.Lock()
	defer c.Unlock()
	if c.ClassIDList == nil {
		c.ClassIDList = make(map[int]int)
	}
	for _, id := range ClassID {
		c.ClassIDList[id] = id
	}
}

func (c *Course) DelClass(ClassID ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range ClassID {
		delete(c.ClassIDList, id)
	}
}

func (c *Course) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}
