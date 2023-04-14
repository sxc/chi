package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/sxc/oishifood/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email     string
		CSRFField template.HTML
	}
	data.Email = r.FormValue("email")
	data.CSRFField = csrf.TemplateField(r)
	u.Templates.New.Execute(w, r, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form submission. Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "<p>Email: %s<p>", r.FormValue("email"))
	fmt.Fprintf(w, "<p>Password: %s</p>", r.FormValue("password"))

	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println()
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		// TODO: Handle error
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	setCookie(w, CookieSession, session.Token)

	// http.SetCookie(w, cookie)
	http.Redirect(w, r, "/users/me", http.StatusFound)
	// fmt.Fprintf(w, "User created: %v", user)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println()
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	// Add this code AFTER a user is successfully authenticated, but BEFORE the
	// any writes to the ResponseWriter.
	// Create a new session

	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	// Add a cookie to the ResponseWriter
	setCookie(w, CookieSession, session.Token)

	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	// Get the session cookie
	// tokenCookie, err := r.Cookie("session")
	token, err := readCookie(r, CookieSession)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	// Get the user from the session token
	user, err := u.SessionService.User(token)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Current user: %s\n", user)

	// fmt.Fprintf(w, "The email cookie is set to: %s", user)
	// fmt.Fprintf(w, "Headers: %+v\n", r.Header)
}

func (u Users) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	// Get the session cookie
	// tokenCookie, err := r.Cookie("session")
	token, err := readCookie(r, CookieSession)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	// // Get the user from the session token
	// tok, err := u.SessionService.User(token)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Redirect(w, r, "/signin", http.StatusFound)
	// 	return
	// }

	// Delete the session
	err = u.SessionService.Delete(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	// Delete the cookie
	// setCookie(w, CookieSession, "")
	deleteCookie(w, CookieSession)
	http.Redirect(w, r, "/signin", http.StatusFound)
}
