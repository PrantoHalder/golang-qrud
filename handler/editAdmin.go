package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"main.go/storage"
)

func (h Handler) EditAdmin(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	editUser, err := h.storage.GetAdminByID(id)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	var form AdminFrom
	form.User = *editUser
	form.CSRFToken = nosurf.Token(r)
	PareseEditAdminTemplate(w, form)
}
func (h Handler) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
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
		PareseEditAdminTemplate(w, AdminFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}

	_, err1 := h.storage.UpdateAdmin(user)
	if err1 != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/users/showadmin", http.StatusSeeOther)
}

func PareseEditAdminTemplate(w http.ResponseWriter, data any) {
	
	
	t, err := template.ParseFiles("assets/templates/admin/adminEdit.html")
	if err != nil {
		log.Fatalf("%v", err)
	}

	t.Execute(w, data)
}

func(h Handler) PareseEditAdminTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("adminEdit.html")
	if t == nil {
		log.Fatal("can not look up adminEdit.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up adminEdit.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}