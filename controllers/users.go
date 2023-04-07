package controllers

import (
	"fmt"
	"net/http"

	"github.com/sxc/oishifood/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Something went wrong.", http.StatusInternalServerError)
	// 	return
	// }
	fmt.Fprint(w, "Email", r.PostForm.Get("email"))
	fmt.Fprint(w, "Password", r.PostForm.Get("password"))
	fmt.Fprint(w, "Test")

	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println()
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %v", user)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data)
}
