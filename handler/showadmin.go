package handler

import (
	"html/template"
	"log"
	"net/http"

	"main.go/storage"
)


func (h Handler) ShowAdmin (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	st := r.FormValue("SearchTerm")
	uf := storage.AdminFilter{
		SearchTerm: st,
	}
	ListUser,err :=h.storage.ListAdmin(uf)
	if err != nil {
		http.Error(w,"Internal Server error",http.StatusInternalServerError)
	}
	data := AdminList{
		Users:      ListUser,
		SearchTerm: st,
	}
	ParseShowAdminTemplate(w,data)
}

func ParseShowAdminTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/showadmin.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)
}

func(h Handler) ParseShowAdminTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("showadmin.html")
	if t == nil {
		log.Fatal("can not look up showadmin.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up showadmin.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
