package main

import (
	_ "CourseManagement/web"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
