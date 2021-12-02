package Http

import (
	"net/http"
	"newstudentmanage/dto"
)

type HttpDBConnect struct {
	dto.DBFile
	Mux *http.ServeMux
}

func (hdbc *HttpDBConnect) StartConnect(filename string, url string) {
	hdbc.FileName = filename
	hdbc.Mux = http.NewServeMux()
	http.ListenAndServe(url, hdbc.Mux)
	hdbc.Mux.Handle("/delete", http.HandlerFunc(hdbc.Delete))
	hdbc.Mux.Handle("/insert", http.HandlerFunc(hdbc.Insert))
}
