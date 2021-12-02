package newstudentmanage

import (
	"fmt"
	"newstudentmanage/dto"
	"testing"
)

func Test1(t *testing.T) {
	var d = dto.DBFile{}
	d.DBMO = dto.Student{}
	d.DBMOS = &dto.StuList{}
	d.FileName = "students.json"
	d.RealInsert(dto.Student{Id: "2017213119", Name: "Rong Bailin"})
	d.RealInsert(dto.Student{Id: "2017213119", Name: "Rong Bailin"})
	d.RealInsert(dto.Student{Id: "2017213119", Name: "Rong Bailin"})
	d.RealInsert(dto.Student{Id: "2017213119", Name: "Rong Bailin"})

	err := d.WriteFile()
	fmt.Println(err)
}
