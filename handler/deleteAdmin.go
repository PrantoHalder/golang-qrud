package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (h Handler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	err := h.storage.DeleteAdminByID(id)
	if err != nil {
		log.Printf("This error is inside DeleteAdmin Handler after chi %#v",err)
		http.Redirect(w,r,"/internalservererror",http.StatusSeeOther)
	}
	http.Redirect(w,r,"/users/showadmin",http.StatusSeeOther)
}
