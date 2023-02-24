package handler

import (
	"log"
	"net/http"
	"text/template"
)


func (h Handler) ShowClass (w http.ResponseWriter, r *http.Request){
	
	ListUser,err :=h.storage.ListClass()
	if err != nil {
		http.Error(w,"Internal Server error",http.StatusInternalServerError)
	}
	ParseShowClassTemplate(w,ListUser)
}

func ParseShowClassTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/showclass.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)
}