package handler

import (
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/form"
	"main.go/storage"
)

type Handler struct {
	sessionManager *scs.SessionManager
	decoder        *form.Decoder
	storage        dbstorage
	Templates      *template.Template
	staticFiles    fs.FS
	templateFiles  fs.FS
}
type ErrorPage struct {
	Code    int
	Message string
}

type dbstorage interface {
	//show contents
	ListUser(storage.UserFilter) ([]storage.User, error)
	ListAdmin(storage.AdminFilter) ([]storage.Admin, error)
	ListFaculty(storage.AdminFilter) ([]storage.Admin, error)
	ListClass() ([]storage.Class, error)
	ListUserResult(uf storage.ResultFilter) ([]storage.User, error)
	//Create contents
	CreateUser(u storage.User) (*storage.User, error)
	CreateAdmin(u storage.Admin) (*storage.Admin, error)
	CreateFaculty(u storage.Admin) (*storage.Admin, error)
	CreateClass(u storage.Class) (*storage.Class, error)
	CreateSubject(u storage.Subject) (*storage.Subject, error)

	// insert contents
	InsertMark(s storage.StudentSubject) (*storage.StudentSubject, error)

	//update contents
	UpdateUser(u storage.User) (*storage.User, error)
	UpdateAdmin(u storage.Admin) (*storage.Admin, error)
	UpdateFaculty(u storage.Admin) (*storage.Admin, error)
	UpdateClass(u storage.Class) (*storage.Class, error)
	UpdateSubject(u storage.Subject) (*storage.Subject, error)
	UpdateMarksbyID(marks string, id string) error
	UpdateStatus(status bool, id int) error
	UpdateAdminStatus(status bool, id int) error
	UpdatefacultyStatus(status bool, id int) error

	//accessing contents
	GetUserByID(string) (*storage.User, error)
	GetAdminByID(id string) (*storage.Admin, error)
	GetFacultyByID(id string) (*storage.Admin, error)
	GetClassByID(string) (*storage.Class, error)
	GetClass() ([]storage.Class, error)
	GetSubjectByClassID(id string) ([]storage.Subject, error)
	GetSubjectByID(id string) (*storage.Subject, error)
	GetSubjectByClassID2(class int) ([]storage.Subject, error)
	GetStudent() ([]storage.User, error)
	GetSubject(id string) ([]storage.SubjectFrom, error)
	GetMarkByID(id string) ([]storage.MarkIN, error)
	MarkEdit(id string) (*storage.MarkEdit, error)
	GetStatusbyUsernameQuery(username string) ([]storage.User, error)
	GetFacultyByUsername(username string) (*storage.Admin, error)
	GetStatusbyUsernameQueryOFFaculty(username string) ([]storage.User, error)
	GetStatusbyUsernameQueryOFUsers(username string) ([]storage.User, error)
	GetUsersByUsername(username string) (*storage.Admin, error)
	GetFacultyEdit(id string) ([]storage.User, error)
	GetUserEditQuery(id string) ([]storage.User, error)

	//login with different roles
	GetAdminByUsername(username string) (*storage.Admin, error)

	//Delete contents
	DeleteUserByID(id string) error
	DeleteAdminByID(id string) error
	DeleteFacultyByID(id string) error
	DeleteClassyByID(id string) error
	DeleteSubjectyByID(id string) error

	//marks input
	Markcreate(s storage.StudentSubject) (*storage.StudentSubject, error)

	//status edit
	StatusEdit(id string) ([]storage.User, error)
	StatusEditAdmin(id string) ([]storage.User, error)
	StatusEditfaculty(id string) ([]storage.User, error)
}

func NewHandler(sm *scs.SessionManager, formDecoder *form.Decoder, storage dbstorage, staticFiles, templateFiles fs.FS) *chi.Mux {
	h := &Handler{
		sessionManager: sm,
		decoder:        formDecoder,
		storage:        storage,
		staticFiles:    staticFiles,
		templateFiles:  templateFiles,
	}
	h.ParseTemplates()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Get("/mainhome", h.MainHome)
		r.Get("/createadmin", h.AdminCreateOutside)
		r.Post("/outsideadminstore", h.AdminStoreOutside)
		r.Get("/future", h.Future)
		r.Get("/career", h.Career)
		r.Get("/faculty", h.Faculty)
		r.Get("/departments", h.Departemnts)
		r.Get("/login", h.Login)
		r.Post("/login", h.LoginPostHandler)
	})

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(h.staticFiles))))

	r.Route("/users", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/home", h.Home)
		r.Get("/inputmarks", h.InputMarks)
		r.Get("/insertmarks", h.InsertMarks)
		r.Get("/show", h.Show)
		r.Get("/showadmin", h.ShowAdmin)
		r.Get("/showfaculty", h.ShowFaculty)
		r.Get("/showclass", h.ShowClass)
		r.Get("/showresult", h.ShowResult)
		r.Get("/showsubjectdetails", h.ShowSubjectDetails)
		r.Get("/showsubject", h.ShowSubject)
		r.Get("/create", h.Create)
		r.Post("/store", h.Store)
		r.Get("/admincreate", h.AdminCreate)
		r.Get("/createsubject", h.SubjectCreate)
		r.Post("/adminstore", h.AdminStore)
		r.Get("/facultycreate", h.CreateFaculty)
		r.Get("/addclass", h.AddClass)
		r.Post("/facultystore", h.FacultyStore)
		r.Post("/storesubject", h.SubjectStore)
		r.Post("/storeclass", h.StoreClass)
		r.Post("/storemarks", h.StoreMarks)
		r.Get("/{id:[0-9]+}/edit", h.Edit)
		r.Get("/{id:[0-9]+}/editmarks", h.EditMark)
		r.Get("/{id:[0-9]+}/editclass", h.EditClass)
		r.Get("/{id:[0-9]+}/editadmin", h.EditAdmin)
		r.Get("/{id:[0-9]+}/editsubject", h.EditSubject)
		r.Get("/{id:[0-9]+}/editstatusadmin", h.EditAdminStatus)
		r.Get("/{id:[0-9]+}/editstatusfaculty", h.EditFacultyStatus)
		r.Get("/{id:[0-9]+}/statusedit", h.EditStatus)
		r.Get("/{id:[0-9]+}/editclass", h.EditClass)
		r.Post("/{id:[0-9]+}/update", h.Update)
		r.Post("/{id:[0-9]+}/updateclass", h.UpdateClass)
		r.Post("/{id:[0-9]+}/updatemarks", h.UpdateMark)
		r.Get("/{id:[0-9]+}/editfaculty", h.EditFaculty)
		r.Post("/{id:[0-9]+}/updatefaculty", h.UpdateFaculty)
		r.Post("/{id:[0-9]+}/updateadmin", h.UpdateAdmin)
		r.Post("/{id:[0-9]+}/updatesubject", h.UpdateSubject)
		r.Get("/{id:[0-9]+}/delete", h.Delete)
		r.Get("/{id:[0-9]+}/deleteadmin", h.DeleteAdmin)
		r.Get("/{id:[0-9]+}/deletefaculty", h.DeleteFaculty)
		r.Get("/{id:[0-9]+}/deleteclass", h.DeleteClass)
		r.Get("/{id:[0-9]+}/deletesubject", h.DeleteSubject)
		r.Get("/{id:[0-9]+}/showresultdetails", h.ShowResultDetails)
	})
	r.Route("/facultys", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/{id:[0-9]+}/home", h.FacultyHome)
		r.Get("/inputmarks", h.InputMarks)
	})
	r.Route("/students", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/{id:[0-9]+}/home", h.StudentsHome)

	})
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/logout", h.Logouthandler)
		r.Get("/logoutfaculty", h.LogoutFacultyhandler)
		r.Get("/logoutstudents", h.LogoutStudentshandler)
	})

	return r
}
func (h Handler) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := h.sessionManager.GetString(r.Context(), "userID")
		uID, err := strconv.Atoi(userID)
		if err != nil {
			log.Fatalln(err)
		}
		if uID <= 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func (h *Handler) ParseTemplates() error {
	templates := template.New("web-templates").Funcs(template.FuncMap{
		"calculatePreviousPage": func(currentPageNumber int) int {
			if currentPageNumber == 1 {
				return 0
			}

			return currentPageNumber - 1
		},

		"calculateNextPage": func(currentPageNumber, totalPage int) int {
			if currentPageNumber == totalPage {
				return 0
			}

			return currentPageNumber + 1
		},
	}).Funcs(sprig.FuncMap())

	tmpl := template.Must(templates.ParseFS(h.templateFiles, "*/*.html"))
	if tmpl == nil {
		log.Fatalln("unable to parse templates")
	}

	h.Templates = tmpl
	return nil
}
