package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Student struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

func (stu Student) GetDBMOId() string {
	return stu.Id
}
func (stu Student) GetDBMOName() string {
	return stu.Name
}

func (Student) AnalyseJsonFile(filename string, dbmos *DBMOS) {
	jsonFile, _ := os.Open(filename)
	defer jsonFile.Close()

	var students []Student
	decoder := json.NewDecoder(jsonFile)
	decoder.Decode(&students)

	for index := 0; index < len(students); index++ {
		(*dbmos).RealInsert(students[index])
	}
}

func (Student) AnalyseHttpInsert(req *http.Request) DBMO {
	var stu Student
	stu.Id = req.Form["Id"][1]
	stu.Name = req.Form["Name"][1]
	return stu
}

func (Student) AnalyseHttpDelete(req *http.Request) string {
	return req.Form["Id"][1]
}

type StuList []DBMO

func (list *StuList) RealInsert(dbmo DBMO) {
	*list = append(*list, dbmo)
}

func (list StuList) RealDelete(studentId string) []DBMO {
	for index := 0; index < len(list); index++ {
		if list[index].GetDBMOId() == studentId {
			copy(list[index:], list[index+1:])
			break
		}
	}
	return list[:len(list)-1]
}

func (list StuList) WriteFile(filename string) error {
	var students []Student

	for index := 0; index < len(list); index++ {
		stu := list[index].(Student)
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

func (list StuList) GetList() []DBMO {
	return list
}
