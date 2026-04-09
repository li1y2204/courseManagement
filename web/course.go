package web

import (
	"CourseManagement/data/class"
	"CourseManagement/data/course"
	"CourseManagement/data/stu"
	"CourseManagement/data/teacher"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AddCourse(w http.ResponseWriter, req *http.Request) {
	Name := req.FormValue("name")
	teacherIds := req.FormValue("teacher_Ids")
	var teacherIdList []int = make([]int, 0)
	json.Unmarshal([]byte(teacherIds), &teacherIdList)

	c := &course.Course{
		Name:        Name,
		StuIDList:   make(map[int]int),
		ClassIDList: make(map[int]int),
		LastClassID: 0,
	}
	c.AddTeacher(teacherIdList...)
	_, err := courseData.Add(c)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, c.String())
}

func GetCourseList(w http.ResponseWriter, req *http.Request) {
	list, err := courseData.Get()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	str, err := json.Marshal(list)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(str))
}

// 课程开班
func NewClassByCourse(w http.ResponseWriter, req *http.Request) {
	courseId, _ := strconv.Atoi(req.FormValue("course_id"))
	className := req.FormValue("class_name")
	headmasterID, _ := strconv.Atoi(req.FormValue("head_master_id"))
	if courseId == 0 {
		err := errors.New("课程ID为0")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	courseList, err := courseData.Get(courseId)
	if err != nil || len(courseList) == 0 {
		err := errors.New("课程信息为空")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	courseObj := courseList[0]

	c := &class.Class{
		Name:         className,
		StuIDList:    make(map[int]int, 0),
		HeadmasterID: headmasterID,
		CourseID:     courseId,
	}
	classID, err := classData.Add(c)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	courseObj.LastClassID = classID
	courseObj.AddClass(classID)
	err = courseData.Edit(courseObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	var t *teacher.Teacher
	teacherList, err := teacherData.Get(headmasterID)
	if err == nil {
		if len(teacherList) > 0 {
			t = teacherList[0]
		} else {
			t = &teacher.Teacher{}
		}
	} else {
		log.Println(err)
	}
	mp := make(map[string]interface{})
	mp["CourseID"] = courseObj.ID
	mp["CourseName"] = courseObj.Name
	mp["ClassID"] = c.ID
	mp["ClassName"] = c.Name
	mp["HeadmasterID"] = c.HeadmasterID
	mp["HeadmasterName"] = t.Name
	res, err := json.Marshal(mp)
	fmt.Fprintf(w, string(res))
}

// 课程报名
func SignUpCourse(w http.ResponseWriter, req *http.Request) {
	stuId, _ := strconv.Atoi(req.FormValue("stu_id"))
	courseId, _ := strconv.Atoi(req.FormValue("course_id"))
	if stuId == 0 {
		err := errors.New("请输入学员信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	if courseId == 0 {
		err := errors.New("请输入课程信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	stuObjList, err := stuData.Get(stuId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var stuObj *stu.Stu
	if len(stuObjList) > 0 {
		stuObj = stuObjList[0]
	}
	if stuObj == nil {
		err := errors.New("未找到学员信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	courseObjList, err := courseData.Get(courseId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var courseObj *course.Course
	if len(courseObjList) > 0 {
		courseObj = courseObjList[0]
	}
	if courseObj == nil {
		err := errors.New("未找到课程信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	classList, err := classData.Get(courseObj.LastClassID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var classObj *class.Class
	if len(classList) > 0 {
		classObj = classList[0]
	}
	if classObj == nil {
		err := errors.New("课程需要先开班")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	stuObj.AddCourse(courseObj.ID)
	stuObj.AddClass(classObj.ID)
	classObj.AddStu(stuObj.ID)
	courseObj.AddStu(stuObj.ID)
	err = stuData.Edit(stuObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	err = classData.Edit(classObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	err = courseData.Edit(courseObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	mp := map[string]interface{}{
		"StuName":    stuObj.Name,
		"CourseName": courseObj.Name,
		"ClassName":  classObj.Name,
	}
	res, _ := json.Marshal(mp)
	fmt.Fprintf(w, string(res))
}

// 退出课程
func SignOutCourse(w http.ResponseWriter, req *http.Request) {
	//找到要删除的stu和course的信息
	stuId, _ := strconv.Atoi(req.FormValue("stu_id"))
	courseId, _ := strconv.Atoi(req.FormValue("course_id"))
	if stuId == 0 {
		err := errors.New("请输入学员信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	if courseId == 0 {
		err := errors.New("请输入课程信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	stuObjList, err := stuData.Get(stuId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var stuObj *stu.Stu
	if len(stuObjList) > 0 {
		stuObj = stuObjList[0]
	}
	if stuObj == nil {
		err := errors.New("未找到学员信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	courseObjList, err := courseData.Get(courseId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	var courseObj *course.Course
	if len(courseObjList) > 0 {
		courseObj = courseObjList[0]
	}
	if courseObj == nil {
		err := errors.New("未找到课程信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	stuObj.DelCourse(courseObj.ID)
	courseObj.DelStu(stuObj.ID)
	classIdList := make([]int, len(courseObj.ClassIDList)) //创造一个新的来存储要删除的course里边的班级id列表
	i := 0
	for _, classId := range courseObj.ClassIDList {
		classIdList[i] = classId
		i++
	}
	classNameList := make([]string, len(courseObj.ClassIDList))
	classObjList, err := classData.Get(classIdList...) //用得到的需要删除的course中的班级id列表来get，id所属的班级信息列表
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	j := 0
	for _, classObj := range classObjList { //在所得到的班级信息列表中查询，
		stuObj.DelClass(classObj.ID)     //学生删除对应班级信息
		classObj.DelStu(stuObj.ID)       //班级删除对应学生id
		classData.Edit(classObj)         //更新class表
		classNameList[j] = classObj.Name //表示以及更新完成该class的信息，可供待会输出哪些班级做了修改
		j++
	}
	stuData.Edit(stuObj)
	courseData.Edit(courseObj)
	mp := make(map[string]interface{})
	mp["StuName"] = stuObj.Name
	mp["CourseName"] = courseObj.Name
	mp["Class"] = strings.Join(classNameList, ",")
	res, _ := json.Marshal(mp)
	fmt.Fprintf(w, string(res))
}

type CourseResp struct {
	CourseID    int                `json:"course_id"`
	CourseName  string             `json:"course_name"`
	StuCount    int                `json:"stu_count"`
	TeacherList []*teacher.Teacher `json:"teacher_list"`
	StuList     []*stu.Stu         `json:"stu_list"`
}

func GetHotCourseList(w http.ResponseWriter, req *http.Request) {
	list, err := courseData.Get()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	teacherIds := make([]int, 0)
	stuIds := make([]int, 0)
	for _, l := range list { //遍历每一门课程，l 代表当前这门课程对象 (Course)
		for key, _ := range l.TeacherIDList {
			teacherIds = append(teacherIds, key) //将 map 的键 (教师 ID) 追加到 teacherIds 切片中
		}
		for key, _ := range l.StuIDList {
			stuIds = append(stuIds, key)
		}
	}
	teacherList, err := teacherData.Get(teacherIds...) // 它会一次性从数据库/内存中捞出所有指定 ID 的教师对象
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	teacherMap := make(map[int]*teacher.Teacher) //初始化一个 Map ，键 (Key): int (教师 ID) ，值 (Value): *teacher.Teacher (教师对象指针) ，这个 Map 将作为“快速查找表”
	for _, t := range teacherList {              //遍历列表，填入 Map
		teacherMap[t.ID] = t //将教师对象以 "ID -> 对象" 的形式存入 Map，例如：teacherMap[101] = 指向 ID 为 101 的教师对象的指针
	}
	stuList, err := stuData.Get(stuIds...)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	stuMap := make(map[int]*stu.Stu)
	for _, s := range stuList {
		stuMap[s.ID] = s
	}

	res := make([]*CourseResp, 0)
	for _, l := range list { //遍历每一门课程，l 代表当前这门课程对象 (Course)
		cr := &CourseResp{
			CourseID:    l.ID,
			CourseName:  l.Name,
			StuCount:    len(l.StuIDList),
			TeacherList: make([]*teacher.Teacher, len(l.TeacherIDList)),
			StuList:     make([]*stu.Stu, len(l.StuIDList)),
		}
		i := 0
		for _, tid := range l.TeacherIDList {
			cr.TeacherList[i] = teacherMap[tid]
			i++
		}
		j := 0
		for _, sid := range l.StuIDList {
			cr.StuList[j] = stuMap[sid]
			j++
		}
		res = append(res, cr)
	}
	res = BubbleSortCompare(res, compare)
	str, _ := json.Marshal(res)
	fmt.Fprintf(w, string(str))
}

func BubbleSortCompare(list []*CourseResp, compareFun func(i, j *CourseResp) int) []*CourseResp {
	length := len(list)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if compareFun(list[j], list[j+1]) == 1 {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func compare(i, j *CourseResp) int {
	if i.StuCount < j.StuCount {
		return 1
	} else {
		return 0
	}
}
