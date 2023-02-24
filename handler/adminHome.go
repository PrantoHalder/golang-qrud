package handler

import (
	"log"
	"net/http"

)

func (h Handler) Home (w http.ResponseWriter, r *http.Request){
	t := h.Templates.Lookup("adminHome.html")
	if t == nil {
		log.Fatal("can not look up adminHome.html")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal("can not look up adminHome.html")
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	
}