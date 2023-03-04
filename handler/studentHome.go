package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (h Handler) StudentsHome (w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	User,err := h.storage.GetUserEditQuery(id)
	if err != nil {
		log.Println(err)
	}
	h.ParseStudentHomeTemplate(w,User)
}
func (h Handler) ParseStudentHomeTemplate(w http.ResponseWriter,data any){
	t := h.Templates.Lookup("studentHome.html")
	if t == nil {
		log.Fatal("can not look up studentHome.html")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up studentHome.html")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}