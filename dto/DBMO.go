package dto

import (
	"net/http"
)

type DBMO interface {
	AnalyseHttpInsert(req *http.Request) DBMO
	AnalyseHttpDelete(req *http.Request) string
	AnalyseJsonFile(string, *DBMOS)
	GetDBMOId() string
}

type DBMOS interface {
	RealInsert(DBMO)
	RealDelete(Id string) []DBMO
	GetList() []DBMO
	WriteFile(filename string) error
}

type DBFile struct {
	DBMO
	DBMOS
	FileName string
}

func (dbf DBFile) ReadFile() {

	switch dbf.FileName {
	case "student.json":
		dbf.FileName = "student.json"
		dbf.DBMOS = new(TeaList)
		dbf.DBMO = new(Student)
	default:
	}
	dbf.AnalyseJsonFile(dbf.FileName, &dbf.DBMOS)
}

func (dbf DBFile) WriteFile() error {
	err := dbf.DBMOS.WriteFile(dbf.FileName)
	if err != nil {
		return err
	}
	return nil
}

func (dbf DBFile) Insert(w http.ResponseWriter, req *http.Request) {

	dbf.ReadFile()
	InsObj := dbf.AnalyseHttpInsert(req)
	dbf.RealInsert(InsObj)
	dbf.WriteFile()

}

func (dbf DBFile) Delete(w http.ResponseWriter, req *http.Request) {

	dbf.ReadFile()
	DeleteId := dbf.AnalyseHttpDelete(req)
	dbf.RealDelete(DeleteId)
	dbf.WriteFile()

}
