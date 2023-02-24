package handler

import (
	"log"
	"net/http"
)


func(h Handler)Future(w http.ResponseWriter, r *http.Request){
	h.ParseFutureTemplate(w,nil)
}
func(h Handler) ParseFutureTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("future.html")
	if t == nil {
		log.Fatal("can not look up future.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up future.html template")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}
