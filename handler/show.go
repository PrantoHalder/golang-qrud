package handler

import (
	"html/template"
	"log"
	"net/http"

	"main.go/storage"
)



type UserList struct {
	Users      []storage.User
	SearchTerm string
}
type AdminList struct {
	Users      []storage.Admin
	SearchTerm string
	
}

func (h Handler) Show(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	
	st := r.FormValue("SearchTerm")
	uf := storage.UserFilter{
		SearchTerm: st,
	}
	ListUser, err := h.storage.ListUser(uf)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}

	data := UserList{
		Users:      ListUser,
		SearchTerm: st,
	}
	ParseShowTemplate(w, data)
}

func ParseShowTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/show.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, data)
}

func (h Handler) ParseShowTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("show.html")
	if t == nil {
		log.Fatal("can not look up show.html template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up show.html template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
