package controllers

import (
	"net/http"

	"github.com/sxc/oishifood/models"
)

type Users struct {
	Templates struct {
		New Template
	}
	UserService models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
