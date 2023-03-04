package handler

import (
	"log"
	"net/http"
)


func(h Handler)Logouthandler(w http.ResponseWriter, r *http.Request){
	if err := h.sessionManager.Destroy(r.Context());err!=nil{
		log.Fatal(err)
	}
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}
func(h Handler)LogoutFacultyhandler(w http.ResponseWriter, r *http.Request){
	if err := h.sessionManager.Destroy(r.Context());err!=nil{
		log.Fatal(err)
	}
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}
func(h Handler)LogoutStudentshandler(w http.ResponseWriter, r *http.Request){
	if err := h.sessionManager.Destroy(r.Context());err!=nil{
		log.Fatal(err)
	}
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}