package main

import (
	"lanshan_homework/go1.19.2/go_homework/class_3_work_lv1/api"
	"lanshan_homework/go1.19.2/go_homework/class_3_work_lv1/dao"
)

func main() {
	dao.Start()
	api.InitRouter()
}
