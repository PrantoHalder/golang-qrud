package handler

import (
	"log"
	"net/http"
	"text/template"

	"main.go/storage"
)


func (h Handler) ShowFaculty (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	st := r.FormValue("SearchTerm")
	uf := storage.AdminFilter{
		SearchTerm: st,
	}
	ListUser,err := h.storage.ListFaculty(uf)
	if err != nil {
		http.Error(w,"Internal Server error",http.StatusInternalServerError)
	}
	data := AdminList{
		Users:      ListUser,
		SearchTerm: st,
	}
	ParseShowFacultyTemplate(w,data)
}

func ParseShowFacultyTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/showfaculty.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)
}