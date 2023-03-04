package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (h Handler) FacultyHome (w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	User,err:=h.storage.GetFacultyEdit(id)
	if err != nil {
		http.Error(w,"internal server error", http.StatusInternalServerError)
	}
	fmt.Printf("%#v",User)
	h.ParseFacultyHomeTemplate(w,User)

	
}


func (h Handler) ParseFacultyHomeTemplate(w http.ResponseWriter,data any){
	t := h.Templates.Lookup("facultyHome.html")
	if t == nil {
		log.Fatal("can not look up facultyHome.html")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up facultyHome.html")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}