package handler

import (
	"log"
	"net/http"
	"text/template"
	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
)

func (h Handler) Edit(w http.ResponseWriter, r *http.Request) {
	classlist, err := h.storage.GetClass()
	if err != nil {
		log.Printf("This error is inside Edit Handler after GetClass query %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	id := chi.URLParam(r,"id")
	editUser,err :=h.storage.GetUserByID(id)
	if err !=nil{
		log.Printf("This error is inside Edit Handler after GetUserByID query %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	
	PareseEditUserTemplate(w,UserFrom{
		User:      *editUser,
		Classlist: classlist,
		FormError: map[string]error{},
		CSRFToken: nosurf.Token(r),
	})
}

func PareseEditUserTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/edit.html")
	if err != nil {
		log.Fatalf("%v", err)
	}

	t.Execute(w, data)
}
