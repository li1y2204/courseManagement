package web

import (
	"CourseManagement/data/class"
	"CourseManagement/data/course"
	"CourseManagement/data/stu"
	"CourseManagement/data/teacher"
)

var (
	stuData     = stu.NewStuData()
	teacherData = teacher.NewTeacherData()
	courseData  = course.NewCourseData()
	classData   = class.NewClassData()
)
