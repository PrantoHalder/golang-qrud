package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	err := h.storage.DeleteUserByID(id)
	if err != nil {
		log.Printf("This error is inside Delete Handler after chi %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	http.Redirect(w,r,"/users/show",http.StatusSeeOther)
}
