package web

import (
	"CourseManagement/common"
	"CourseManagement/data/teacher"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func AddTeacher(w http.ResponseWriter, req *http.Request) {
	Name := req.FormValue("name")
	GenderStr := req.FormValue("gender")
	BirthdayStr := req.FormValue("birthday")
	/*log.Printf("[DEBUG] 收到请求 -> Name: [%s], GenderStr: [%s], Birthday: [%s]", Name, GenderStr, BirthdayStr)*/ //测试是否收到数据
	gender, _ := strconv.Atoi(GenderStr)
	t := &teacher.Teacher{
		Name:   Name,
		Gender: common.Gender(gender),
	}
	birthday, err := common.StrToTime(BirthdayStr)
	if err == nil {
		t.Birthday = birthday
	}
	_, err = teacherData.Add(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	fmt.Fprintf(w, t.String())
}

func GetTeacherList(w http.ResponseWriter, req *http.Request) {
	list, err := teacherData.Get()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	bytes, err := json.Marshal(list) //有疑问
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(bytes))
}
