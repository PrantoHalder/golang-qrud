package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
	"main.go/storage"
)
type MarkIn struct{
	MarkIn []storage.MarkIN
	FormError map[string]error
	CSRFToken string
}
type MarkEdit struct{
	MarkIn storage.MarkEdit
	FormError map[string]error
	CSRFToken string
}



func (h Handler) MarksHandler(w http.ResponseWriter, r *http.Request, class int, studentID int) error {

	subject, err := h.storage.GetSubjectByClassID2(class)
	fmt.Println(subject)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	
	for _, s := range subject {
		b := storage.StudentSubject{
			StudentID: studentID,
			SubjectID: s.ID,
			Marks:     0,
		}

		_, err := h.storage.InsertMark(b)
		if err != nil {
			log.Fatalf("%v", err)
			return err
		}
	}
	return nil
}
func (h Handler) ShowResultDetails(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	mark,err := h.storage.GetMarkByID(id)
	if err != nil {
		log.Println(err)
	}
	ParseMarkInTemplate(w,MarkIn{
		MarkIn:    mark,
		FormError: map[string]error{},
		CSRFToken: nosurf.Token(r),
	})

}
func (h Handler) EditMark(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	mark,err := h.storage.MarkEdit(id)
	if err != nil {
		log.Println(mark)
	}
	fmt.Printf("%#v",mark)
	ParseMarkEditTemplate(w,MarkEdit{
		MarkIn:    *mark,
		FormError: map[string]error{},
		CSRFToken: nosurf.Token(r),
	})

}
func (h Handler) UpdateMark(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	var form MarkEdit
	user := storage.MarkEdit{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatalln(err)
	}
	form.MarkIn = user
	fmt.Printf("%#v",user)
	err1 := h.storage.UpdateMarksbyID(user.Marks,id)
	if err1 != nil {
        log.Println(err)
	}
	http.Redirect(w, r, "/users/showresult", http.StatusSeeOther)
}

func ParseMarkInTemplate(w http.ResponseWriter,data any){
	t,err := template.ParseFiles("assets/templates/admin/markIn.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w,data)
}
func ParseMarkEditTemplate(w http.ResponseWriter,data any){
	t,err := template.ParseFiles("assets/templates/admin/editMark.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w,data)
}