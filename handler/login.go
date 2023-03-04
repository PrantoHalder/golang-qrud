package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"golang.org/x/crypto/bcrypt"
	"main.go/storage/postgres"
)

type LoginUser struct {
	Username  string `form:"Username"`
	Password  string
	Loginas   []string
	FormError map[string]error
	CSRFToken string
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.pareseLoginTemplate(w, LoginUser{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	var lf LoginUser
	if err := h.decoder.Decode(&lf, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	if lf.Loginas == nil {
		if err := lf.validate(); err != nil {
			if vErr, ok := err.(validation.Errors); ok {
				lf.FormError = vErr
			}
			h.pareseLoginTemplate(w, LoginUser{
				Username:  lf.Username,
				Password:  "",
				FormError: lf.FormError,
				CSRFToken: nosurf.Token(r),
			})
			return
		}
	}
	for _, value := range lf.Loginas {
		if value == "Admin" {
			user, err := h.storage.GetStatusbyUsernameQuery(lf.Username)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			for _, usr := range user {
				if usr.Status {
					if err := lf.validate(); err != nil {
						if vErr, ok := err.(validation.Errors); ok {
							lf.FormError = vErr
						}
						h.pareseLoginTemplate(w, LoginUser{
							Username:  lf.Username,
							Password:  "",
							FormError: lf.FormError,
							CSRFToken: nosurf.Token(r),
						})
						return
					}
					user, err := h.storage.GetAdminByUsername(lf.Username)
					if err != nil {
						if err.Error() == postgres.NotFound {
							formErr := make(map[string]error)
							formErr["Username"] = fmt.Errorf("credentials does not match")
							lf.FormError = formErr
							lf.CSRFToken = nosurf.Token(r)
							lf.Password = ""
							h.pareseLoginTemplate(w, lf)
							return
						}

						http.Error(w, "internal server error", http.StatusInternalServerError)
						return
					}
					if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lf.Password)); err != nil {
						formErr := make(map[string]error)
						formErr["Username"] = fmt.Errorf("credentials does not match")
						lf.FormError = formErr
						h.pareseLoginTemplate(w, LoginUser{
							Username:  lf.Username,
							Password:  "",
							FormError: lf.FormError,
							CSRFToken: nosurf.Token(r),
						})
						return
					}

					h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(user.ID))
					http.Redirect(w, r, "/users/home", http.StatusSeeOther)
					return
				} else {
					h.pareseInactiveTemplate(w, nil)
				}
			}
		}
	}


	for _, value := range lf.Loginas {
		fmt.Println("inside faculty")
		if value == "Faculty" {
			user, err := h.storage.GetStatusbyUsernameQueryOFFaculty(lf.Username)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			for _, usr := range user {
				if usr.Status {
					if err := lf.validate(); err != nil {
						if vErr, ok := err.(validation.Errors); ok {
							lf.FormError = vErr
						}
						h.pareseLoginTemplate(w, LoginUser{
							Username:  lf.Username,
							Password:  "",
							FormError: lf.FormError,
							CSRFToken: nosurf.Token(r),
						})
						return
					}
					user, err := h.storage.GetFacultyByUsername(lf.Username)
					if err != nil {
						if err.Error() == postgres.NotFound {
							formErr := make(map[string]error)
							formErr["Username"] = fmt.Errorf("credentials does not match")
							lf.FormError = formErr
							lf.CSRFToken = nosurf.Token(r)
							lf.Password = ""
							h.pareseLoginTemplate(w, lf)
							return
						}

						http.Error(w, "internal server error", http.StatusInternalServerError)
						return
					}
					if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lf.Password)); err != nil {
						formErr := make(map[string]error)
						formErr["Username"] = fmt.Errorf("credentials does not match")
						lf.FormError = formErr
						h.pareseLoginTemplate(w, LoginUser{
							Username:  lf.Username,
							Password:  "",
							FormError: lf.FormError,
							CSRFToken: nosurf.Token(r),
						})
						return
					}

					h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(user.ID))
					http.Redirect(w, r,fmt.Sprintf("/facultys/%v/home",user.ID), http.StatusSeeOther)
					return
				} else {
					h.pareseInactiveTemplate(w, nil)
				}
			}
		}
	}


	for _, value := range lf.Loginas {
		fmt.Println("inside faculty")
		if value == "Student" {
			user, err := h.storage.GetStatusbyUsernameQueryOFUsers(lf.Username)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			for _, usr := range user {
				if usr.Status {
					if err := lf.validate(); err != nil {
						if vErr, ok := err.(validation.Errors); ok {
							lf.FormError = vErr
						}
						h.pareseLoginTemplate(w, LoginUser{
							Username:  lf.Username,
							Password:  "",
							FormError: lf.FormError,
							CSRFToken: nosurf.Token(r),
						})
						return
					}
					user, err := h.storage.GetUsersByUsername(lf.Username)
					if err != nil {
						if err.Error() == postgres.NotFound {
							formErr := make(map[string]error)
							formErr["Username"] = fmt.Errorf("credentials does not match")
							lf.FormError = formErr
							lf.CSRFToken = nosurf.Token(r)
							lf.Password = ""
							h.pareseLoginTemplate(w, lf)
							return
						}

						http.Error(w, "internal server error", http.StatusInternalServerError)
						return
					}
					if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lf.Password)); err != nil {
						formErr := make(map[string]error)
						formErr["Username"] = fmt.Errorf("credentials does not match")
						lf.FormError = formErr
						h.pareseLoginTemplate(w, LoginUser{
							Username:  lf.Username,
							Password:  "",
							FormError: lf.FormError,
							CSRFToken: nosurf.Token(r),
						})
						return
					}

					h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(user.ID))
					http.Redirect(w, r, fmt.Sprintf("/students/%v/home",user.ID), http.StatusSeeOther)
					return
				} else {
					h.pareseInactiveTemplate(w, nil)
				}
			}
		}
	}

	

	
	h.pareseLoginTemplate(w, nil)
}

func (h Handler) pareseLoginTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/open/login.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	t.Execute(w, data)
}
func pareseStudentHomeTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/students/studentHome.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.Execute(w, data)
}

func (h Handler) pareseInactiveTemplate(w http.ResponseWriter, data any) {
	t, err := template.ParseFiles("assets/templates/open/Unactive.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t.Execute(w, data)
}

func (lu LoginUser) validate() error {
	return validation.ValidateStruct(&lu, validation.Field(&lu.Username,
		validation.Required.Error("username can not be blank"),
	),
		validation.Field(&lu.Password,
			validation.Required.Error("password can not be blank"),
		),
		validation.Field(&lu.Loginas,
			validation.Required.Error("login role can not be blank"),
		),
	)
}
