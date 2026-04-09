package web

import "net/http"

func init() {
	http.HandleFunc("/add/teacher", AddTeacher)
	http.HandleFunc("/get/teacher", GetTeacherList)

	http.HandleFunc("/add/course", AddCourse)
	http.HandleFunc("/get/course", GetCourseList)

	http.HandleFunc("/add/stu", AddStu)
	http.HandleFunc("/get/stu", GetStuList)

	http.HandleFunc("/course/new/class", NewClassByCourse)
	http.HandleFunc("/sign/up/course", SignUpCourse)
	http.HandleFunc("/sign/out/course", SignOutCourse)
	http.HandleFunc("/get/hot/course/list", GetHotCourseList)

}
