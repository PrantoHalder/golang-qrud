package handler

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)


func(h Handler) Create (w http.ResponseWriter, r *http.Request){
	classlist, err := h.storage.GetClass()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	h.ParseCreateTemplate(w, UserFrom{
		Classlist:     classlist,
		CSRFToken: nosurf.Token(r),
	})
}
func (h Handler) ParseCreateTemplate(w http.ResponseWriter, data any) {
	t,err := template.ParseFiles("assets/templates/admin/create.html")
	if err != nil {
		http.Error(w, "internal server error",http.StatusInternalServerError)
	}
	
	t.Execute(w,data)
}

