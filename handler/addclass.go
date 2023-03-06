package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"main.go/storage"
)

type ClassFrom struct {
	Class     storage.Class
	FormError map[string]error
	CSRFToken string
}

func (h Handler) AddClass(w http.ResponseWriter, r *http.Request) {
	h.ParseAddClassTemplate(w, ClassFrom{
		FormError: map[string]error{},
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) StoreClass(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("This error is inside Store class Handler after parse template %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	form := ClassFrom{}
	user := storage.Class{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Printf("This error is inside Store class Handler after decoding %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	form.Class = user
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseAddClassTemplate(w, ClassFrom{
			Class:     storage.Class{},
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}

	_, err := h.storage.CreateClass(user)
	if err != nil {
		log.Printf("This error is inside Store class Handler CreateClass query %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}

	http.Redirect(w, r, fmt.Sprintln("/users/showclass"), http.StatusSeeOther)

}

func (h Handler) EditClass(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	editUser, err := h.storage.GetClassByID(id)
	if err != nil {
		log.Printf("This error is inside Edit class Handler after GetClassBYID query %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	var form ClassFrom
	form.Class = *editUser
	form.CSRFToken = nosurf.Token(r)
	h.ParseEditClassTemplate(w, form)
}

func (h Handler) UpdateClass(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("This error is inside Update class Handler after getting the id form Chi route %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	fmt.Println(uID)
	class := storage.Class{ID: uID}
	if err := h.decoder.Decode(&class, r.PostForm); err != nil {
		log.Printf("This error is inside Update class Handler after decoding %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	var form ClassFrom
	form.Class = class
	if err := class.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			form.FormError = vErr
		}
		h.ParseEditClassTemplate(w, ClassFrom{
			Class:      class,
			FormError: form.FormError,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
	

	_, err1 := h.storage.UpdateClass(class)
	if err1 != nil {
		log.Printf("This error is inside Update class Handler after Update Class query %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}

	http.Redirect(w, r, "/users/showclass", http.StatusSeeOther)
}


func (h Handler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	err := h.storage.DeleteClassyByID(id)
	if err != nil {
		log.Printf("This error is inside Delete class Handler after getting the id form Chi route %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	http.Redirect(w,r,"/users/showclass",http.StatusSeeOther)
}

func (h Handler) InternalServerError(w http.ResponseWriter, r *http.Request) {
	 h.ParseInternalServerErrorTempate(w,nil)
}


func(h Handler) ParseAddClassTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("addclass.html")
	if t == nil {
		log.Fatal("can not look up addClassTemplate")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal("can not look up addClassTemplate")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}

func (h Handler)ParseEditClassTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("editclass.html")
	if t == nil {
		log.Fatal("can not look up addClassTemplate")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal("can not look up addClassTemplate")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	
}

func (h Handler) ParseInternalServerErrorTempate(w http.ResponseWriter,data any){
	t := h.Templates.Lookup("httpServererror.html")
	if t == nil {
		log.Fatal("can not look up httpServererror")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal("can not look up httpServererror")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
