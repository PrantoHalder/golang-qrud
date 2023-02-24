package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"main.go/storage"
)


func(h Handler) AdminCreate (w http.ResponseWriter, r *http.Request){
	h.ParseAdminCreateTemplate(w, UserFrom{
		CSRFToken: nosurf.Token(r),
	})
}
func(h Handler) AdminStore (w http.ResponseWriter, r *http.Request){
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
		h.ParseAdminCreateTemplate(w,AdminFrom{
			User:      user,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
    _,err := h.storage.CreateAdmin(user)
	if err != nil {
		log.Fatalln(err)
	}

	http.Redirect(w, r, fmt.Sprintln("/users/showadmin"), http.StatusSeeOther)
}
func(h Handler) ParseAdminCreateTemplate(w http.ResponseWriter, data any) {
	t,err := template.ParseFiles("assets/templates/admin/adminSignUp.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	
}
}