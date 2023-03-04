package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
	"main.go/storage"
)
type MarkForm struct {
	SubjectList []storage.SubjectFrom
	FormError map[string]error
	CSRFToken string
}

func (h Handler) InputMarks(w http.ResponseWriter, r *http.Request) {

	studentList, err := h.storage.GetStudent()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	h.ParseInputMarksTemplate(w, UserFrom{
		StudentList: studentList,
		FormError:   map[string]error{},
		CSRFToken:   nosurf.Token(r),
	})
}

func (h Handler) ParseInputMarksTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/admin/markInput.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)
}

func (h Handler) InsertMarks(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	
	Student_id := r.FormValue("Student_id")

    
	subjectList,err := h.storage.GetSubject(Student_id)
	if err != nil{
		log.Println(err)
	}
	fmt.Printf("%#v",subjectList)
	h.ParsrInsertMarksTemplate(w,MarkForm{
		SubjectList: subjectList,
		CSRFToken:   nosurf.Token(r),
	})

}
func (h Handler) ParsrInsertMarksTemplate(w http.ResponseWriter, data any) {
     t,err := template.ParseFiles("assets/templates/admin/insertmarktemplate.html")
	 if err != nil{
		log.Fatalln(err)
	 }
	 t.Execute(w,data)
}

func (h Handler) StoreMarks(w http.ResponseWriter,r *http.Request){
   if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	user := storage.StudentSubject{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Fatal(err)
	}
	for id , mark := range user.Mark{
        p:=storage.StudentSubject{
            ID:id,
			Marks : mark,
		} 
		_,err := h.storage.Markcreate(p)
		if err != nil {
			log.Fatalln(err)
		}
	}
	http.Redirect(w, r, fmt.Sprintln("/users/showresult"), http.StatusSeeOther)
}
func (h Handler)ShowResult(w http.ResponseWriter, r *http.Request){
    if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	
	st := r.FormValue("SearchTerm")
	uf := storage.ResultFilter{
		SearchTerm: st,
	}
	ListUser, err := h.storage.ListUserResult(uf)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}

	data := UserList{
		Users:      ListUser,
		SearchTerm: st,
	}
	h.ParseShowResultTemplate(w,data)
}
func (h Handler) ParseShowResultTemplate(w http.ResponseWriter,data any){
	t,err := template.ParseFiles("assets/templates/admin/showresult.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w,data)
}
