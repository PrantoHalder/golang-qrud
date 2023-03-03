package handler

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
)

func(h Handler) EditStatus (w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	StatusEdit,err := h.storage.StatusEdit(id)
	fmt.Printf("%v",StatusEdit)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	for _,value := range StatusEdit{
         if value.Status {
			value.Status = false
			err := h.storage.UpdateStatus(value.Status,value.ID)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		 }
	}
	for _,value := range StatusEdit{
		if !value.Status {
		   value.Status = true
		   err := h.storage.UpdateStatus(value.Status,value.ID)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}
   }
   http.Redirect(w,r,"/users/show",http.StatusSeeOther)
}
