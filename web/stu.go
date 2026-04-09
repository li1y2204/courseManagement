package web

import (
	"CourseManagement/common"
	"CourseManagement/data/stu"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func AddStu(w http.ResponseWriter, req *http.Request) {
	Name := req.FormValue("name")
	GenderStr := req.FormValue("gender")
	BirthdayStr := req.FormValue("birthday")
	/*log.Printf("[DEBUG] 收到请求 -> Name: [%s], GenderStr: [%s], Birthday: [%s]", Name, GenderStr, BirthdayStr)*/ //测试是否收到数据
	gender, _ := strconv.Atoi(GenderStr)
	s := &stu.Stu{
		Name:   Name,
		Gender: common.Gender(gender),
	}
	birthday, err := common.StrToTime(BirthdayStr)
	if err == nil {
		s.Birthday = birthday
	}
	_, err = stuData.Add(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	fmt.Fprintf(w, s.String())
}

func GetStuList(w http.ResponseWriter, req *http.Request) {
	list, err := stuData.Get()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	bytes, err := json.Marshal(list) //有疑问为什么还要进行序列化
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(bytes))
}
