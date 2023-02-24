package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"main.go/storage"
)


func(h Handler) CreateFaculty (w http.ResponseWriter, r *http.Request){
	ParseFacultyCreateTemplate(w, UserFrom{
		FormError: map[string]error{},
		CSRFToken: nosurf.Token(r),
	})
}
func ParseFacultyCreateTemplate(w http.ResponseWriter, data any) {
	t,err := template.ParseFiles("assets/templates/admin/createfaculty.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, data)
}
//lookup template
func(h Handler) ParseFacultyCreateTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("createfaculty.html")
	if t == nil {
		log.Fatal("can not look up createfaculty.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up createfaculty.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}

func (h Handler) FacultyStore(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
    form := AdminFrom{}
	user := storage.Admin{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
	}
    form.User = user
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		ParseFacultyCreateTemplate(w,AdminFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
    _,err := h.storage.CreateFaculty(user)
	if err != nil {
		log.Fatalln(err)
	}

	http.Redirect(w, r, fmt.Sprintln("/users/showfaculty"), http.StatusSeeOther)
}


func (h Handler) EditFaculty(w http.ResponseWriter, r *http.Request) {
	
	id := chi.URLParam(r,"id")
	editUser,err :=h.storage.GetFacultyByID(id)
	if err !=nil{
		log.Fatalln(err)
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	var form AdminFrom
	form.User = *editUser
    form.CSRFToken =nosurf.Token(r)
	PareseEditFacultyTemplate(w, form)
}

func (h Handler) UpdateFaculty(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	var form AdminFrom
	user := storage.Admin{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	form.User = user
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		PareseEditFacultyTemplate(w,AdminFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}

	_, err1 := h.storage.UpdateFaculty(user)
	if err1 != nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w,r,"/users/showfaculty",http.StatusSeeOther)
}

func (h Handler) DeleteFaculty (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	err := h.storage.DeleteFacultyByID(id)
	if err != nil {
		http.Error(w,"internal serval error",http.StatusInternalServerError)
	}
	http.Redirect(w,r,"/users/showfaculty",http.StatusSeeOther)
}

func PareseEditFacultyTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/editfaculty.html")
	if err != nil {
		log.Fatalf("%v", err)
	}

	t.Execute(w, data)
}


