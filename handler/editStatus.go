package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func(h Handler) EditStatus (w http.ResponseWriter, r *http.Request){
	status := chi.URLParam(r,"status")
	fmt.Println(status)
}