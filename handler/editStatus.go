package handler

import (
	"net/http"
	"github.com/go-chi/chi"
)

func(h Handler) EditStatus (w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	StatusEdit,err := h.storage.StatusEdit(id)
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
func(h Handler) EditAdminStatus (w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	StatusEdit,err := h.storage.StatusEditAdmin(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	for _,value := range StatusEdit{
         if value.Status {
			value.Status = false
			err := h.storage.UpdateAdminStatus(value.Status,value.ID)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		 }
	}
	for _,value := range StatusEdit{
		if !value.Status {
		   value.Status = true
		   err := h.storage.UpdateAdminStatus(value.Status,value.ID)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}
   }
   http.Redirect(w,r,"/users/showadmin",http.StatusSeeOther)
}
func(h Handler) EditFacultyStatus (w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	StatusEdit,err := h.storage.StatusEditfaculty(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	for _,value := range StatusEdit{
         if value.Status {
			value.Status = false
			err := h.storage.UpdatefacultyStatus(value.Status,value.ID)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		 }
	}
	for _,value := range StatusEdit{
		if !value.Status {
		   value.Status = true
		   err := h.storage.UpdatefacultyStatus(value.Status,value.ID)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}
   }
   http.Redirect(w,r,"/users/showfaculty",http.StatusSeeOther)
}
