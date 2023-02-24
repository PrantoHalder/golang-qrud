package handler

import (
	"fmt"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"main.go/storage"
)
type UserFrom struct{
	User storage.User
	Classlist []storage.Class
	StudentList []storage.User
	FormError map[string]error
	CSRFToken string
}
type SubjectFrom struct{
	User storage.Subject
	Classlist []storage.Class
	FormError map[string]error
	CSRFToken string
}
type AdminFrom struct{
	User storage.Admin
	Classlist []storage.Class
	Class storage.Class
	FormError map[string]error
	CSRFToken string
}

func (h Handler) Store(w http.ResponseWriter, r *http.Request) {
	classlist, err := h.storage.GetClass()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
    form := UserFrom{}
	user := storage.User{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
	}
    form.User = user
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseCreateTemplate(w,UserFrom{
			User:      user,
			Classlist: classlist,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
    data,err1 := h.storage.CreateUser(user)
	if err1 != nil {
		log.Fatalln(err)
	}
    if err := h.MarksHandler(w,r, user.Class_id,data.ID);err != nil{
		http.Error(w,"internal error",http.StatusInternalServerError)
	}
	http.Redirect(w, r, fmt.Sprintln("/users/show"), http.StatusSeeOther)
}

