package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Teacher struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

func (tea Teacher) GetDBMOId() string {
	return tea.Id
}
func (tea Teacher) GetDBMOName() string {
	return tea.Name
}

func (Teacher) AnalyseJsonFile(filename string, dbmos *DBMOS) {
	jsonFile, _ := os.Open(filename)
	defer jsonFile.Close()

	var students []Teacher
	decoder := json.NewDecoder(jsonFile)
	decoder.Decode(&students)

	for index := 0; index < len(students); index++ {
		(*dbmos).RealInsert(students[index])
	}
}

func (Teacher) AnalyseHttpInsert(req *http.Request) DBMO {
	var stu Teacher
	stu.Id = req.Form["Id"][1]
	stu.Name = req.Form["Name"][1]
	return stu
}

func (Teacher) AnalyseHttpDelete(req *http.Request) string {
	return req.Form["Id"][1]
}

type TeaList []DBMO

func (list *TeaList) RealInsert(dbmo DBMO) {
	*list = append(*list, dbmo)
}

func (list TeaList) RealDelete(studentId string) []DBMO {
	for index := 0; index < len(list); index++ {
		if list[index].GetDBMOId() == studentId {
			copy(list[index:], list[index+1:])
			break
		}
	}
	return list[:len(list)-1]
}

func (list TeaList) WriteFile(filename string) error {
	var students []Teacher

	for index := 0; index < len(list); index++ {
		stu := list[index].(Teacher)
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

func (list TeaList) GetList() []DBMO {
	return list
}

