package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"main.go/storage"
)

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	classlist, err := h.storage.GetClass()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	form := UserFrom{}
	user := storage.User{}
	user = storage.User{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	form.User = user
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		PareseEditUserTemplate(w,UserFrom{
			User:      user,
			Classlist: classlist,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
	
	

	_, err1 := h.storage.UpdateUser(user)
	if err1 != nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w,r,"/users/show",http.StatusSeeOther)
}
