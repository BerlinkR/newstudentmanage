package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Course struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

func (cou Course) GetDBMOId() string {
	return cou.Id
}
func (cou Course) GetDBMOName() string {
	return cou.Name
}

func (Course) AnalyseJsonFile(filename string, dbmos *DBMOS) {
	jsonFile, _ := os.Open(filename)
	defer jsonFile.Close()

	var students []Course
	decoder := json.NewDecoder(jsonFile)
	decoder.Decode(&students)

	for index := 0; index < len(students); index++ {
		(*dbmos).RealInsert(students[index])
	}
}

func (Course) AnalyseHttpInsert(req *http.Request) DBMO {
	var stu Course
	stu.Id = req.Form["Id"][1]
	stu.Name = req.Form["Name"][1]
	return stu
}

func (Course) AnalyseHttpDelete(req *http.Request) string {
	return req.Form["Id"][1]
}

type CouList []DBMO

func (list *CouList) RealInsert(dbmo DBMO) {
	*list = append(*list, dbmo)
}

func (list CouList) RealDelete(studentId string) []DBMO {
	for index := 0; index < len(list); index++ {
		if list[index].GetDBMOId() == studentId {
			copy(list[index:], list[index+1:])
			break
		}
	}
	return list[:len(list)-1]
}

func (list CouList) WriteFile(filename string) error {
	var students []Course

	for index := 0; index < len(list); index++ {
		stu := list[index].(Course)
		students = append(students, stu)
	}
	fmt.Println(students)
	data, err := json.MarshalIndent(students, "", " ")
	fmt.Println(string(data))
	if err != nil {
		return err
	}
	err2 := ioutil.WriteFile(filename, data, 644)
	if err2 != nil {
		return err2
	}
	return nil

}

func (list CouList) GetList() []DBMO {
	return list
}
