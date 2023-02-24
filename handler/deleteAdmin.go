package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h Handler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	err := h.storage.DeleteAdminByID(id)
	if err != nil {
		http.Error(w,"internal serval error",http.StatusInternalServerError)
	}
	http.Redirect(w,r,"/users/showadmin",http.StatusSeeOther)
}
