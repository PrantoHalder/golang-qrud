package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"main.go/storage"
)

type ClassForm struct {
	Student   []storage.Class
	CSRFToken string
	FormError map[string]error
}
type SubjectForm struct {
	ClassForm []storage.Class
	Student   storage.Subject
	CSRFToken string
	FormError map[string]error
}
type SubjectList struct {
	ID []storage.Subject
}

func (h Handler) SubjectCreate(w http.ResponseWriter, r *http.Request) {
	classlist, err := h.storage.GetClass()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	ParseCreateSubjectTemplate(w, SubjectForm{
		ClassForm: classlist,
		Student:   storage.Subject{},
		CSRFToken: nosurf.Token(r),
	})
}
func (h Handler) ShowSubject(w http.ResponseWriter, r *http.Request) {
	classlist, err := h.storage.GetClass()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	user := storage.Subject{}
	form := SubjectFrom{}
	
	ParseShowSubjectTemplate(w, SubjectFrom{
		User:      user,
		Classlist: classlist,
		FormError: form.FormError,
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) ShowSubjectDetails(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
   class_id := r.FormValue("Class")
   fmt.Println(class_id)
   subject,err := h.storage.GetSubjectByClassID(class_id)
   if err != nil {
	http.Error(w,"internal server error",http.StatusInternalServerError)
   }
   data := SubjectList{
   	ID: subject,
   }
   ParseShowSubjectDetailsTemplate(w,data)
}

func (h Handler) SubjectStore(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	form := SubjectForm{}
	user := storage.Subject{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
	}

	classlist, err := h.storage.GetClass()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	form.Student = user

	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		ParseCreateSubjectTemplate(w, SubjectForm{
			Student:   user,
			ClassForm: classlist,
			CSRFToken: nosurf.Token(r),
			FormError: form.FormError,
		})
		return
	}
	_, err1 := h.storage.CreateSubject(user)
	if err1 != nil {
		log.Fatalln(err)
	}

	http.Redirect(w, r, fmt.Sprintln("/users/showsubject"), http.StatusSeeOther)
}
func (h Handler) EditSubject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	editUser, err := h.storage.GetSubjectByID(id)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	var form SubjectFrom
	form.User = *editUser
	form.CSRFToken = nosurf.Token(r)
	ParseEditSubjectTemplate(w, form)
}
func (h Handler) UpdateSubject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	var form SubjectFrom
	user := storage.Subject{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	form.User = user
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		ParseEditSubjectTemplate(w, SubjectFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
	_, err1 := h.storage.UpdateSubject(user)
	if err1 != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/users/showsubject", http.StatusSeeOther)
}

func (h Handler) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.storage.DeleteSubjectyByID(id)
	if err != nil {
		http.Error(w, "internal serval error", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/users/showsubject", http.StatusSeeOther)
}

func ParseCreateSubjectTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/subjectcreate.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, data)
}
func ParseShowSubjectTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/showSubject.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, data)
}
func ParseShowSubjectDetailsTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/showSubjectDetails.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, data)
}
func ParseEditSubjectTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/editsubject.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, data)
}
