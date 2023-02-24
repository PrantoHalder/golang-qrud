package handler

import (
	"log"
	"net/http"
)

func(h Handler) Faculty(w http.ResponseWriter, r *http.Request){
		t := h.Templates.Lookup("faculty.html")
		if t == nil {
			log.Fatal("can not look up faculty.html template")
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		}
		if err := t.Execute(w, nil); err != nil {
			log.Fatal("can not look up faculty.html template")
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		}
	
}